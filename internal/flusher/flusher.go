package flusher

//go:generate mockgen -destination=../mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-roadmap-api/internal/flusher Flusher

import (
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/ozoncp/ocp-roadmap-api/internal/repo"
	"log"
)

type Flusher interface {
	Flush(entities []entity.Roadmap) []entity.Roadmap
}

type Flush struct {
	repo repo.Repo
}

func NewFlush(repo repo.Repo) *Flush {
	return &Flush{
		repo: repo,
	}
}

func (f *Flush) Flush(entities []entity.Roadmap) []entity.Roadmap {
	if e := f.repo.AddEntities(entities); e != nil {
		log.Fatalf("error while add entities, err: %s\n", e)
	}

	return make([]entity.Roadmap, 0)
}
