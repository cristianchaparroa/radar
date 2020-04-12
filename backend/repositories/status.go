package repositories

import (
	"radar/domain"
)

type Status interface {
	Create(location domain.Location,statusName string) (*domain.Status, error)
	FindLastByUserID(profileID string) (domain.Status, error)
	FindAllByUserID(userID string) ([]domain.Status, error)
	InactivateLast(profileID string) error
}