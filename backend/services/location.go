package services

import (
	"github.com/google/uuid"
	"radar/domain"
	"radar/providers/sql"
	"radar/repositories"
	"time"
)

type ILocation interface {
	RegisterLocation(l domain.Location) (domain.Location, error)
}

type Location struct {
	sql                sql.Client
	locationRepository repositories.Location
}

func NewLocation(client sql.Client) ILocation {
	service := repositories.NewLocation(client)
	return &Location{sql: client, locationRepository: service}

}

func (s *Location) RegisterLocation(l domain.Location) (domain.Location, error) {

	t := time.Now()
	l.UpdatedAt = t
	l.CreateAt = t
	l.ID = uuid.New().String()

	return l, s.locationRepository.Register(l)
}
