package presenters

import "time"

type Location struct {
	ID        string    `json:"id:"`
	ProfileID string    `json:"profileID"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updateAt"`
}
