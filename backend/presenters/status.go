package presenters

import "time"

type Status struct {
	ID string `json:"id"`
	ProfileID string     `json:"profile_id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}
