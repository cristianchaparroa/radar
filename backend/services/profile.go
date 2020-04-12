package services

import (
	"radar/domain"
	"radar/providers/sql"
	"radar/repositories"
	"time"
)

type IProfile interface {
	Create(p domain.Profile) (domain.Profile, error)
	GetProfile(id string) (*domain.Profile, error)
}

type Profile struct {
	sql               sql.Client
	profileRepository repositories.IProfile
}

func NewProfile(sql sql.Client) IProfile {
	pr := repositories.NewProfile(sql)
	return &Profile{sql: sql, profileRepository: pr}
}

func (s *Profile) Create(p domain.Profile) (domain.Profile, error) {
	now := time.Now()
	p.CreatedAt = &now
	return p, s.profileRepository.Create(p)
}

func (s *Profile) GetProfile(id string) (*domain.Profile, error) {
	return s.profileRepository.FindByID(id)
}
