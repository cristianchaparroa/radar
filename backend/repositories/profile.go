package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"radar/adapters"
	"radar/domain"
	"radar/models"
	"radar/providers/sql"
)

type IProfile interface {
	FindByID(id string) (*domain.Profile, error)
	Create(p domain.Profile) error
}

type Profile struct {
	db *gorm.DB
}

func NewProfile(sql sql.Client) IProfile {
	return &Profile{db: sql.GetConnection()}
}

func (r *Profile) FindByID(id string) (*domain.Profile, error) {

	fmt.Println("--> FindByID", id)
	var profile models.Profile

	err := r.db.Where("id=?", id).First(&profile).Error

	if err != nil {
		return nil, err
	}
	p := adapters.MapProfileEntityToDomain(profile)
	return &p, nil
}

func (r *Profile) Create(p domain.Profile) error {
	profile := adapters.MapProfileDomainToEntity(p)
	return r.db.Save(&profile).Error
}
