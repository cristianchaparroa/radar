package repositories

import (
	"radar/domain"
)

// Infections retrieves all the data related to the possible infections between
// profiles that have reported the statuses.
type Infections interface {

	// It retrieves the confirmed positive cases
	GetAllPositiveCases() ([]domain.CurrentProfile, error)

	// GetContacts retrieves the possible contacts with another profiles.
	GetContacts(positives []domain.CurrentProfile) ([]domain.Profile, error)
}