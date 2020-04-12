package domain

import (
	"time"
)

const (
	PositiveStatus  = "POSITIVE"
	NegativeStatus  = "NEGATIVE"
	RecoveredStatus = "RECOVERED"
)

type Status struct {
	ID         string
	ProfileID  string
	Name       string
	CreatedAt  *time.Time
	Current 	bool
	Location Location
}