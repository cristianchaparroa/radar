package domain

import "time"

type Location struct {
	ID        string
	UserID    string
	Latitude  float64
	Longitude float64
	CreateAt  time.Time
	UpdatedAt time.Time
}

