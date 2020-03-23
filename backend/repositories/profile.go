package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"radar/dataprovider/sql"
	"radar/entities"
)

type IProfile interface {
	FindByID(id string) (*entities.Profile, error)
	Create(p entities.Profile) error
}

type Profile struct {
	db *gorm.DB
}

func NewProfile(sql sql.Client) IProfile {
	return &Profile{db: sql.GetConnection()}
}

func (r *Profile) FindByID(id string) (*entities.Profile, error) {

	fmt.Println("--> FindByID", id)
	var profile entities.Profile
	err := r.db.Where("id=?", id).First(&profile).Error

	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *Profile) Create(p entities.Profile) error {
	return r.db.Save(&p).Error
}
