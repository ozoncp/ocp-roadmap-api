package repo

//go:generate mockgen -destination=../mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-roadmap-api/internal/repo Repo

import "github.com/ozoncp/ocp-roadmap-api/internal/entity"

type Repo interface {
	AddEntities(entities []entity.Roadmap) error
	ListEntities(limit, offset uint64) ([]entity.Roadmap, error)
	DescribeEntity(entityId uint64) (*entity.Roadmap, error)
}
