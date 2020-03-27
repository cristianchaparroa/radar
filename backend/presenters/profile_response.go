package presenters

import "time"

type Profile struct {
	ID        string     `json:"id"`
	DeviceID  string     `json:"device_id" `
	CreatedAt *time.Time `json:"created_at"`
}

type ProfileResponse struct {
	Profile Profile `json:"profile"`
	Status  Status  `json:"status"`
}