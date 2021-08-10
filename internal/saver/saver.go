package saver

import (
	"context"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/ozoncp/ocp-roadmap-api/internal/flusher"
	"log"
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
	isClosed bool
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
			case <-ticker.C:
				s.flusher.Flush(s.readChannel())
			case <-s.ctx.Done():
				s.flusher.Flush(s.readChannel())
				return
			}
		}
	}()
}

func (s *save) readChannel() []entity.Roadmap {
	var roadMaps []entity.Roadmap
	for i := 0; i < len(s.buffer); i++ {
		roadMaps = append(roadMaps, <-s.buffer)
	}

	return roadMaps
}

func (s *save) Save(entity entity.Roadmap) {
	if len(s.buffer) == cap(s.buffer) {
		log.Fatalf("channel is full, wait for flush it!\n")
		return
	}

	if s.isClosed {
		log.Fatalf("channel is closed!\n")
		return
	}

	s.buffer <- entity
}

func (s *save) Close() {
	close(s.buffer)
	s.isClosed = true
	s.ctx.Done()
}
