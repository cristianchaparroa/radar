package entities

import "time"

type Profile struct {
	ID        string     `gorm:"primary_key" json:"id"`
	DeviceID  string     `db:"device_id" json:"device_id" `
	CreatedAt *time.Time `json:"created_at"`
}
