package api

import (
	"context"
	ocp_roadmap_api "github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api"
	"github.com/rs/zerolog/log"
)

type RoadmapAPI struct {
	ocp_roadmap_api.UnimplementedOcpRoadmapApiServer
}

func NewRoadmapAPI() ocp_roadmap_api.OcpRoadmapApiServer {
	return &RoadmapAPI{}
}

func (r *RoadmapAPI) CreateRoadmap(ctx context.Context, request *ocp_roadmap_api.CreateRoadmapRequest) (*ocp_roadmap_api.CreateRoadmapResponse, error) {
	log.Info().Msg("Created roadmap")
	return &ocp_roadmap_api.CreateRoadmapResponse{}, nil
}

func (r *RoadmapAPI) DescribeRoadmap(ctx context.Context, request *ocp_roadmap_api.DescribeRoadmapRequest) (*ocp_roadmap_api.DescribeRoadmapResponse, error) {
	log.Info().Msgf("Describe roadmap #%v", request.GetId())
	return &ocp_roadmap_api.DescribeRoadmapResponse{}, nil
}

func (r *RoadmapAPI) ListRoadmap(ctx context.Context, request *ocp_roadmap_api.ListRoadmapRequest) (*ocp_roadmap_api.ListRoadmapResponse, error) {
	log.Info().Msgf("Show list of roadmaps")
	return &ocp_roadmap_api.ListRoadmapResponse{}, nil
}

func (r *RoadmapAPI) RemoveRoadmap(ctx context.Context, request *ocp_roadmap_api.RemoveRoadmapRequest) (*ocp_roadmap_api.RemoveRoadmapResponse, error) {
	log.Info().Msgf("Remove roadmap #%v", request.GetId())
	return &ocp_roadmap_api.RemoveRoadmapResponse{}, nil
}
