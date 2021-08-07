package saver

import (
	"context"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/ozoncp/ocp-roadmap-api/internal/flusher"
	"time"
)

type Saver interface {
	Save(entity entity.Roadmap)
	Init()
	Close()
}

type save struct {
	ctx      context.Context
	tick     time.Duration
	flusher  flusher.Flusher
	entities []entity.Roadmap
	buffer   chan entity.Roadmap
}

func NewSaver(ctx context.Context, flusher flusher.Flusher, tick time.Duration, capacity uint) Saver {
	return &save{
		ctx:     ctx,
		flusher: flusher,
		tick:    tick,
		buffer:  make(chan entity.Roadmap, capacity),
	}
}

func (s *save) Init() {
	go func() {
		ticker := time.NewTicker(s.tick)
		defer ticker.Stop()

		for {
			select {
			case e := <-s.buffer:
				s.entities = append(s.entities, e)
			case <-ticker.C:
				s.entities = s.flusher.Flush(s.entities)
			case <-s.ctx.Done():
				s.entities = s.flusher.Flush(s.entities)
				close(s.buffer)
				return
			}
		}
	}()
}

func (s *save) Save(entity entity.Roadmap) {
	s.buffer <- entity
}

func (s *save) Close() {
	s.ctx.Done()
}
