package main

import (
	"context"
	"os"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/v3"
	"github.com/xds-sotw-server/config"
	"github.com/xds-sotw-server/resources"
	run "github.com/xds-sotw-server/server"
	"github.com/xds-sotw-server/storage"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	logger := Logger{Debug: cfg.Debug}
	// Create a cache
	cache := cache.NewSnapshotCache(false, cache.IDHash{}, logger)

	// Initialize the database
	db := storage.Init(cfg)

	rateLimitStore := storage.NewRateLimitStore(db)
	resource := resources.NewResources(rateLimitStore)
	// Create the snapshot that we'll serve to Envoy
	snapshot := resource.GenerateSnapshot()
	if err := snapshot.Consistent(); err != nil {
		logger.Errorf("Snapshot is inconsistent: %+v\n%+v", snapshot, err)
		os.Exit(1)
	}
	logger.Debugf("Will serve snapshot %+v", snapshot)

	// Add the snapshot to the cache
	if err := cache.SetSnapshot(context.Background(), cfg.NodeId, snapshot); err != nil {
		logger.Errorf("Snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}

	// Run the xDS server
	ctx := context.Background()
	cb := &test.Callbacks{Debug: logger.Debug}
	srv := server.NewServer(ctx, cache, cb)
	run.RunServer(ctx, srv, uint(cfg.Port))
}
