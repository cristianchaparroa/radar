package entities

import "time"

type Location struct {
	ID        string `gorm:"primary_key"`
	UserID    string
	Latitude  float64
	Longitude float64
	CreateAt  time.Time
	UpdatedAt time.Time
}
