package server

import (
	"context"
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xds-sotw-server/models"
	"github.com/xds-sotw-server/resources"
	"github.com/xds-sotw-server/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	pb "github.com/xds-sotw-server/proto"
)

const (
	grpcKeepaliveTime        = 30 * time.Second
	grpcKeepaliveTimeout     = 5 * time.Second
	grpcKeepaliveMinTime     = 30 * time.Second
	grpcMaxConcurrentStreams = 1000000
)

type RateLimitServiceServer struct {
	pb.UnimplementedRateLimitServiceServer
	store                 *storage.RateLimitStore
	configUpdateEventChan chan struct{}
}

// SaveRateLimitDescriptor save rate limit descriptor
func (c *RateLimitServiceServer) SaveRateLimitDescriptor(_ context.Context, in *pb.SaveDescriptorRequest) (*pb.SaveDescriptorResponse, error) {
	rateLimit, err := c.store.GetRateLimitByDomain(in.Domain)
	if err != nil {
		log.Errorf("failed to get rate limit by domain %s: %v", in.Domain, err)
		return &pb.SaveDescriptorResponse{Success: false}, err
	}

	descriptor := &models.RateLimitDescriptor{
		RateLimitId: rateLimit.ID,
		Key:         in.Key,
		Value:       *in.Value,
		Policy: models.RateLimitPolicy{
			Unit:            *in.Policy.Unit,
			RequestsPerUnit: *in.Policy.RequestPerUnit,
		},
		ShadowMode: false,
	}
	if err := c.store.SaveRateLimitDescriptor(descriptor); err != nil {
		log.Errorf("failed to save rate limit descriptor: %v", err)
		return &pb.SaveDescriptorResponse{Success: false}, err
	}
	c.configUpdateEventChan <- struct{}{}
	return &pb.SaveDescriptorResponse{Success: true}, nil
}

// ListRateLimit list all rate limit
func (c *RateLimitServiceServer) ListRateLimit(_ context.Context, _ *pb.ListRateLimitRequest) (*pb.ListRateLimitResponse, error) {
	rateLimits, err := c.store.ListRateLimit()
	if err != nil {
		log.Errorf("failed to list rate limits: %v", err)
		return nil, err
	}

	// TODO: support sub descriptors
	// convert rate limit descriptor
	convert := func(descriptors []models.RateLimitDescriptor) []*pb.RateLimitDescriptor {
		_descriptors := make([]*pb.RateLimitDescriptor, 0, len(descriptors))
		for _, descriptor := range descriptors {
			_descriptors = append(_descriptors, &pb.RateLimitDescriptor{
				Key:   descriptor.Key,
				Value: &descriptor.Value,
				Policy: &pb.RateLimitPolicy{
					Unit:           &descriptor.Policy.Unit,
					RequestPerUnit: &descriptor.Policy.RequestsPerUnit,
				},
			})
		}
		return _descriptors
	}
	rls := make([]*pb.RateLimit, 0, len(rateLimits))
	for _, rl := range rateLimits {
		rls = append(rls, &pb.RateLimit{
			Domain:      rl.Domain,
			Name:        rl.Name,
			Descriptor_: convert(rl.Descriptors),
		})
	}

	return &pb.ListRateLimitResponse{RateLimit: rls}, nil
}

// RunServer starts an xDS server at the given port.
func RunServer(_ context.Context, srv server.Server, resource *resources.Resources, port uint) {
	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    grpcKeepaliveTime,
			Timeout: grpcKeepaliveTimeout,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             grpcKeepaliveMinTime,
			PermitWithoutStream: true,
		}),
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterRateLimitServiceServer(grpcServer, &RateLimitServiceServer{
		store:                 resource.RateLimitStore,
		configUpdateEventChan: resource.ConfigUpdateEventChan,
	})
	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, srv)

	log.Printf("Management server listening on %d", port)
	if err = grpcServer.Serve(lis); err != nil {
		log.Println(err)
	}
}
