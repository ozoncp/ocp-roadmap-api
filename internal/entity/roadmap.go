package entity

import (
	"fmt"
	"time"
)

type Roadmap struct {
	Id        uint64
	UserId    uint64
	Link      string
	CreatedAt time.Time
}

func NewRoadMap(id, userId uint64, link string, createdAt time.Time) *Roadmap {
	return &Roadmap{Id: id, UserId: userId, Link: link, CreatedAt: createdAt}
}

func (r *Roadmap) String() string {
	return fmt.Sprintf("Id: %v, UserId: %v, Link: %v", r.Id, r.UserId, r.Link)
}
