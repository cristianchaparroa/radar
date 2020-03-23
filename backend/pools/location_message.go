package pools

import (
	"github.com/google/uuid"
	"time"
)

type LocationMessage struct {
	ID        string `json:"id"` // ID is a uud identification for the message
	UserID    string
	Latitude  float64
	Longitude float64
	Time      time.Time `json:"time"` // Time in which was sent the message
}

func NewLocationMessage(userID string, latitude, longitude float64) LocationMessage {

	uuid := uuid.New()

	return LocationMessage{
		ID:        uuid.String(),
		Time:      time.Now(),
		Latitude:  latitude,
		Longitude: longitude,
		UserID:    userID,
	}
}
