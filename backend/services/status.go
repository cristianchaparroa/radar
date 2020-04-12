package services

import (
	"radar/domain"
	"radar/providers/sql"
	"radar/repositories"
)

type Status interface {
	Create(location domain.Location, name string) (*domain.Status, error)
	InactivateLast(profileID string) error
}

type Statuses struct {
	sql        sql.Client
	repository repositories.Status
}

func NewStatus(sql sql.Client) Status {
	repository := repositories.NewStatus(sql)
	return &Statuses{sql: sql, repository: repository}
}

func (s *Statuses) Create(l domain.Location,name string) (*domain.Status, error) {

	err := s.InactivateLast(l.ProfileID)

	if err != nil {
		return nil, err
	}

	return s.repository.Create(l,name)
}

func (s *Statuses) InactivateLast(profileID string) error {
	return s.repository.InactivateLast(profileID)
}