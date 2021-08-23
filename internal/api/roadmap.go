package api

import (
	"context"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/ozoncp/ocp-roadmap-api/internal/repo"
	ocp_roadmap_api "github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api"
	"github.com/rs/zerolog/log"
)

type RoadmapAPI struct {
	repository repo.Repo
	ocp_roadmap_api.UnimplementedOcpRoadmapApiServer
}

func NewRoadmapAPI(repo repo.Repo) ocp_roadmap_api.OcpRoadmapApiServer {
	return &RoadmapAPI{repository: repo}
}

func (r *RoadmapAPI) CreateRoadmap(ctx context.Context, request *ocp_roadmap_api.CreateRoadmapRequest) (*ocp_roadmap_api.CreateRoadmapResponse, error) {
	roadMap := []entity.Roadmap{
		*entity.NewRoadMap(
			request.GetRoadmap().GetId(),
			request.GetRoadmap().GetUserId(),
			request.GetRoadmap().GetLink(),
			request.GetRoadmap().GetCreatedAt().AsTime(),
		),
	}

	if err := r.repository.AddEntities(ctx, roadMap); err != nil {
		return &ocp_roadmap_api.CreateRoadmapResponse{}, err
	}
	log.Info().Msg("Created roadmap")
	return &ocp_roadmap_api.CreateRoadmapResponse{}, nil
}

func (r *RoadmapAPI) DescribeRoadmap(ctx context.Context, request *ocp_roadmap_api.DescribeRoadmapRequest) (*ocp_roadmap_api.DescribeRoadmapResponse, error) {
	roadmap, err := r.repository.DescribeEntity(ctx, request.GetId())

	if err != nil {
		return &ocp_roadmap_api.DescribeRoadmapResponse{}, err
	}

	response := &ocp_roadmap_api.DescribeRoadmapResponse{
		Roadmap: &ocp_roadmap_api.Roadmap{
			Id:     roadmap.Id,
			UserId: roadmap.UserId,
			Link:   roadmap.Link,
		},
	}
	return response, nil
}

func (r *RoadmapAPI) ListRoadmap(ctx context.Context, request *ocp_roadmap_api.ListRoadmapRequest) (*ocp_roadmap_api.ListRoadmapResponse, error) {
	result, err := r.repository.ListEntities(ctx, request.GetLimit(), request.GetOffset())
	if err != nil {
		log.Error().Msgf("Error while fetch data, %s", err.Error())
		return &ocp_roadmap_api.ListRoadmapResponse{}, err
	}

	var rmaps []*ocp_roadmap_api.Roadmap
	for _, v := range result {
		item := &ocp_roadmap_api.Roadmap{
			Id:     v.Id,
			UserId: v.UserId,
			Link:   v.Link,
		}
		rmaps = append(rmaps, item)
	}
	response := &ocp_roadmap_api.ListRoadmapResponse{
		Roadmaps: rmaps,
	}

	return response, nil
}

func (r *RoadmapAPI) RemoveRoadmap(ctx context.Context, request *ocp_roadmap_api.RemoveRoadmapRequest) (*ocp_roadmap_api.RemoveRoadmapResponse, error) {
	if err := r.repository.RemoveEntity(ctx, request.GetId()); err != nil {
		return &ocp_roadmap_api.RemoveRoadmapResponse{}, err
	}

	return &ocp_roadmap_api.RemoveRoadmapResponse{Removed: true}, nil
}
