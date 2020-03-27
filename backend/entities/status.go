package entities

import (
	"github.com/google/uuid"
	"time"
)

const (
	PositiveStatus  = "POSITIVE"
	NegativeStatus  = "NEGATIVE"
	RecoveredStatus = "RECOVERED"
)

type Status struct {
	ID        string `gorm:"primary_key" json:"id"`
	ProfileID string
	LocationID string
	Name      string
	CreatedAt *time.Time
}

func NewStatus(userID, locationID, name string) *Status {
	t := time.Now()
	return &Status{
		ID:        uuid.New().String(),
		ProfileID: userID,
		Name:      name,
		LocationID:locationID,
		CreatedAt: &t,
	}
}
