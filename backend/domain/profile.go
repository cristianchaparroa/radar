package domain

import "time"

type Profile struct {
	ID        string
	DeviceID  string
	CreatedAt *time.Time
}
