package repositories

import (
	"github.com/jinzhu/gorm"
	"radar/adapters"
	"radar/domain"
	"radar/providers/sql"
)

type LocationRegister struct {
	db *gorm.DB
}

func NewLocation(client sql.Client) Location {
	return &LocationRegister{db: client.GetConnection()}
}

func (r *LocationRegister) Register(location domain.Location) error {
	l := adapters.MapLocationDomainToEntity(location)
	err := r.db.Save(&l).Error
	return err
}