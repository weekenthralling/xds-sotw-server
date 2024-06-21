package resources

import (
	"context"
	"os"
	"strconv"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	rlsconfig "github.com/envoyproxy/go-control-plane/ratelimit/config/ratelimit/v3"
	"github.com/xds-sotw-server/logger"

	"github.com/xds-sotw-server/config"
	"github.com/xds-sotw-server/models"
	"github.com/xds-sotw-server/storage"
)

type Resources struct {
	Cache                 cache.SnapshotCache
	ConfigUpdateEventChan chan struct{}
	Logger                logger.Logger
	NodeId                string
	RateLimitStore        *storage.RateLimitStore
	Version               *int64
}

func NewResources(cfg *config.Config) *Resources {
	log := logger.Logger{Debug: cfg.Debug}

	var version int64 = 1
	// init db resource
	db := storage.Init(cfg)
	rateLimitStore := storage.NewRateLimitStore(db)

	r := &Resources{
		RateLimitStore:        rateLimitStore,
		Version:               &version,
		ConfigUpdateEventChan: make(chan struct{}),
		NodeId:                cfg.NodeId,
		Logger:                log,
	}

	// init snapshotCache and set default config
	snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, log)

	snapshot := r.GenerateSnapshot(&version)
	if err := snapshotCache.SetSnapshot(context.Background(), cfg.NodeId, snapshot); err != nil {
		log.Errorf("Snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}
	r.Cache = snapshotCache

	// watch resource and update snapshot
	go r.watch()
	return r
}

func (r *Resources) watch() {
	for {
		<-r.ConfigUpdateEventChan
		r.Logger.Infof("Config update event received")

		*r.Version += 1
		// generate new snapshot
		snapshot := r.GenerateSnapshot(r.Version)
		// update cache snapshot
		if err := r.Cache.SetSnapshot(context.Background(), r.NodeId, snapshot); err != nil {
			r.Logger.Infof("Failed to generate snapshot from DB: %v", err)
		}
		r.Logger.Debugf("Generate new snapshot for nodeId: %s, snapshot: %+v, version: %v", r.NodeId, snapshot, *r.Version)
	}
}

// GenerateSnapshot generate snapshot with version
func (r *Resources) GenerateSnapshot(version *int64) *cache.Snapshot {
	rateLimits, _ := r.RateLimitStore.ListRateLimit()
	snap, _ := cache.NewSnapshot(strconv.FormatInt(*version, 10),
		map[resource.Type][]types.Resource{
			resource.RateLimitConfigType: makeRlsConfig(rateLimits),
		},
	)
	return snap
}

// makeRlsConfig converts RateLimit to envoy's RateLimitConfig
func makeRlsConfig(rateLimits []models.RateLimit) []types.Resource {
	var resources []types.Resource
	for _, rateLimit := range rateLimits {
		envoyConfig := &rlsconfig.RateLimitConfig{
			Name:        rateLimit.Name,
			Domain:      rateLimit.Domain,
			Descriptors: convertDescriptors(rateLimit.Descriptors),
		}
		resources = append(resources, envoyConfig)
	}

	return resources
}

// convertDescriptors is a helper function to convert RateLimitDescriptor to envoy's RateLimitDescriptor
func convertDescriptors(descriptors []models.RateLimitDescriptor) []*rlsconfig.RateLimitDescriptor {
	var envoyDescriptors []*rlsconfig.RateLimitDescriptor
	for _, descriptor := range descriptors {
		envoyDescriptor := &rlsconfig.RateLimitDescriptor{
			Key:   descriptor.Key,
			Value: descriptor.Value,
			RateLimit: &rlsconfig.RateLimitPolicy{
				Unit:            rlsconfig.RateLimitUnit(descriptor.Unit),
				RequestsPerUnit: uint32(descriptor.RequestsPerUnit),
			},
			ShadowMode:  descriptor.ShadowMode,
			Descriptors: convertDescriptors(descriptor.Descriptors),
		}
		envoyDescriptors = append(envoyDescriptors, envoyDescriptor)
	}
	return envoyDescriptors
}
