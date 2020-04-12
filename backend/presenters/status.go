package presenters

import "time"

type Status struct {
	ID        string     `json:"id"`
	ProfileID string     `json:"profileID"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
	Current bool `json:"current"`
}
