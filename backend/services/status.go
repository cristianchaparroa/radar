package services

import (
	"radar/entities"
	"radar/providers/sql"
	"radar/repositories"
)

type IStatus interface {
	Create(userID, name string) (*entities.Status, error)
}

type Status struct {
	sql sql.Client
	repository repositories.IStatus
}

func NewStatus(sql sql.Client) IStatus {
	repository := repositories.NewStatus(sql)
	return &Status{sql:sql, repository:repository}
}

func (s *Status) Create(userID, name string ) (*entities.Status, error) {
	return s.repository.Create(userID, name)
}