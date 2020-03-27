package services

import (
	"github.com/google/uuid"
	"radar/entities"
	"radar/providers/sql"
	"radar/repositories"
	"time"
)

type ILocation interface {
	RegisterLocation(l entities.Location)  (entities.Location, error)
}

type Location struct {
	sql                sql.Client
	locationRepository repositories.ILocation
}

func NewLocation(client sql.Client) ILocation {
	service := repositories.NewLocation(client)
	return &Location{sql: client, locationRepository:service}

}

func (s *Location) RegisterLocation(l entities.Location) (entities.Location, error) {

	t := time.Now()
	l.UpdatedAt = t;
	l.CreateAt = t;
	l.ID = uuid.New().String()

	return  l, s.locationRepository.RegisterLocation(l)
}
