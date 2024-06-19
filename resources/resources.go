package resources

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	rls_config "github.com/envoyproxy/go-control-plane/ratelimit/config/ratelimit/v3"
	"github.com/xds-sotw-server/models"
	"github.com/xds-sotw-server/storage"
)

type Resources struct {
	rateLimitStore *storage.RateLimitStore
}

func NewResources(rateLimitStore *storage.RateLimitStore) *Resources {
	return &Resources{
		rateLimitStore: rateLimitStore,
	}
}

func (r Resources) GenerateSnapshot() *cache.Snapshot {
	rateLimits, _ := r.rateLimitStore.ListRateLimit()

	snap, _ := cache.NewSnapshot("1",
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
		envoyConfig := &rls_config.RateLimitConfig{
			Name:        rateLimit.Name,
			Domain:      rateLimit.Domain,
			Descriptors: convertDescriptors(rateLimit.Descriptors),
		}
		resources = append(resources, envoyConfig)
	}

	return resources
}

// convertDescriptors is a helper function to convert RateLimitDescriptor to envoy's RateLimitDescriptor
func convertDescriptors(descriptors []models.RateLimitDescriptor) []*rls_config.RateLimitDescriptor {
	var envoyDescriptors []*rls_config.RateLimitDescriptor
	for _, descriptor := range descriptors {
		envoyDescriptor := &rls_config.RateLimitDescriptor{
			Key:   descriptor.Key,
			Value: descriptor.Value,
			RateLimit: &rls_config.RateLimitPolicy{
				Unit:            rls_config.RateLimitUnit(descriptor.Unit),
				RequestsPerUnit: descriptor.RequestsPerUnit,
			},
			ShadowMode:  descriptor.ShadowMode,
			Descriptors: convertDescriptors(descriptor.Descriptors),
		}
		envoyDescriptors = append(envoyDescriptors, envoyDescriptor)
	}
	return envoyDescriptors
}
