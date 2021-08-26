package flusher

//go:generate mockgen -destination=../mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-roadmap-api/internal/flusher Flusher

import (
	"context"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/ozoncp/ocp-roadmap-api/internal/repo"
	"github.com/ozoncp/ocp-roadmap-api/internal/utils"
	"log"
)

type Flusher interface {
	Flush(entities []entity.Roadmap) []entity.Roadmap
}

type Flush struct {
	chunkSize uint
	repo      repo.Repo
}

func NewFlush(chunkSize uint, repo repo.Repo) *Flush {
	return &Flush{
		chunkSize: chunkSize,
		repo:      repo,
	}
}

func (f *Flush) Flush(entities []entity.Roadmap) []entity.Roadmap {
	chunks := utils.SplitToBulks(entities, f.chunkSize)

	for _, v := range chunks {
		if _, e := f.repo.MultiCreateEntity(context.Background(), v); e != nil {
			log.Fatalf("error while add entities, err: %s\n", e)
		}
	}

	return make([]entity.Roadmap, 0)
}
