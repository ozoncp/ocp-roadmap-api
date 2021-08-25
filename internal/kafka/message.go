package kafka

import "time"

type EventType int

const (
	Create EventType = iota
	Update
	Delete
)

type Message struct {
	Type EventType
	Body body
}

type body struct {
	Id        uint64
	Action    string
	CreatedAt int64
}

func CreateMessage(et EventType, id uint64) Message {
	return Message{
		Type: et,
		Body: body{
			Id:        id,
			Action:    getEventName(et),
			CreatedAt: time.Now().Unix(),
		},
	}
}

func getEventName(et EventType) string {
	switch et {
	case Create:
		return "Created"
	case Update:
		return "Updated"
	case Delete:
		return "Deleted"
	}
	return "undefined type"
}
