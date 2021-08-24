package api

import (
	"context"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/ozoncp/ocp-roadmap-api/internal/repo"
	"github.com/ozoncp/ocp-roadmap-api/internal/utils"
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

func (r *RoadmapAPI) UpdateRoadmap(ctx context.Context, request *ocp_roadmap_api.UpdateRoadmapRequest) (*ocp_roadmap_api.UpdateRoadmapResponse, error) {
	response := &ocp_roadmap_api.UpdateRoadmapResponse{}
	roadmap := *entity.NewRoadMap(request.Id, request.UserId, request.Link, request.CreatedAt.AsTime())
	updated, err := r.repository.UpdateEntity(ctx, roadmap)
	if err != nil {
		log.Error().Msgf("error while update roadmap # %v, err: %v", err, request.Id)
		return response, err
	}

	response.Updated = updated
	return response, nil
}

func (r *RoadmapAPI) MultiCreateRoadmaps(ctx context.Context, request *ocp_roadmap_api.MultiCreateRoadmapRequest) (*ocp_roadmap_api.MultiCreateRoadmapResponse, error) {
	var data []entity.Roadmap

	for _, v := range request.Roadmaps {
		item := entity.NewRoadMap(v.Id, v.UserId, v.Link, v.CreatedAt.AsTime())
		data = append(data, *item)
	}

	response := &ocp_roadmap_api.MultiCreateRoadmapResponse{}
	bulks := utils.SplitToBulks(data, 3)
	for i := 0; i < len(bulks); i++ {
		ids, err := r.repository.MultiCreateEntity(ctx, bulks[i])
		if err != nil {
			log.Error().Msgf("error while multi create roadmap, err: %v", err)
			return response, err
		}

		response.RoadmapsIds = append(response.RoadmapsIds, ids...)
	}

	log.Info().Msg("Multi Create roadmaps")
	return response, nil
}

func (r *RoadmapAPI) CreateRoadmap(ctx context.Context, request *ocp_roadmap_api.CreateRoadmapRequest) (*ocp_roadmap_api.CreateRoadmapResponse, error) {
	roadMap := *entity.NewRoadMap(
		request.GetRoadmap().GetId(),
		request.GetRoadmap().GetUserId(),
		request.GetRoadmap().GetLink(),
		request.GetRoadmap().GetCreatedAt().AsTime(),
	)

	if err := r.repository.CreateEntity(ctx, roadMap); err != nil {
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
