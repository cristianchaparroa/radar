package services

import (
	"radar/entities"
	"radar/providers/sql"
	"radar/repositories"
)

type ILocation interface {
	RegisterLocation(l entities.Location) error
}

type Location struct {
	sql                sql.Client
	locationRepository repositories.ILocation
}

func NewLocation(client sql.Client) ILocation {
	return &Location{sql: client}

}

func (s *Location) RegisterLocation(l entities.Location) error {
	return s.locationRepository.RegisterLocation(l)
}
