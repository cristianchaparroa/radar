package models

import "time"

type Location struct {
	ID        string `gorm:"primary_key"`
	ProfileID string
	Latitude  float64
	Longitude float64
	CreateAt  time.Time
	UpdatedAt time.Time
}
