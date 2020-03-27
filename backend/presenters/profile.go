package presenters

import "time"

type Profile struct {
	ID        string     `json:"id"`
	DeviceID  string     `json:"device_id" `
	CreatedAt *time.Time `json:"created_at"`
}