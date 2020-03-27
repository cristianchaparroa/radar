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
	UserID    string
	Name      string
	CreatedAt *time.Time
}

func NewStatus(userID, name string) *Status {
	t := time.Now()
	return &Status{
		ID:        uuid.New().String(),
		UserID:    userID,
		Name:      name,
		CreatedAt: &t,
	}
}
