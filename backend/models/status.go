package models

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
	ID         string `gorm:"primary_key" json:"id"`
	ProfileID  string

	Name       string
	CreatedAt  *time.Time
	Current 	bool
	Location Location `gorm:"foreignkey:ID"`
}

func NewStatus(l Location, name string) *Status {
	t := time.Now()
	return &Status{
		ID:         uuid.New().String(),
		ProfileID:  l.ProfileID,
		Name:       name,
		CreatedAt:  &t,
		Current:true,
		Location:l,
	}
}
