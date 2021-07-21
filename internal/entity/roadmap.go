package entity

import "fmt"

type Roadmap struct {
	Id     uint64
	UserId uint64
	Link   string
}

func NewRoadMap(id, userId uint64, link string) *Roadmap {
	return &Roadmap{Id: id, UserId: userId, Link: link}
}

func (r *Roadmap) String() string {
	return fmt.Sprintf("Id: %v, UserId: %v, Link: %v", r.Id, r.UserId, r.Link)
}
