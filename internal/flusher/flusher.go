package flusher

import "github.com/ozoncp/ocp-roadmap-api/internal/entity"

type Flusher interface {
	Flush(entities []entity.Roadmap) []entity.Roadmap
}
