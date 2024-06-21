package main

import (
	"context"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/v3"
	"github.com/xds-sotw-server/config"
	"github.com/xds-sotw-server/resources"
	run "github.com/xds-sotw-server/server"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	resource := resources.NewResources(cfg)

	// Run the xDS server
	ctx := context.Background()
	cb := &test.Callbacks{Debug: cfg.Debug}
	srv := server.NewServer(ctx, resource.Cache, cb)
	run.RunServer(ctx, srv, resource, uint(cfg.Port))
}
