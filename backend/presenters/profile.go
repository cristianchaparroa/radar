package presenters

import "time"

type Profile struct {
	ID        string     `json:"id"`
	DeviceID  string     `json:"deviceID" `
	CreatedAt *time.Time `json:"createdAt"`
}
