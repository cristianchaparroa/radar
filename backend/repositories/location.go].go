package repositories

import (
	"github.com/jinzhu/gorm"
	"radar/providers/sql"
	"radar/entities"
)

type ILocation interface {
	RegisterLocation(location entities.Location) error
}

type Location struct {
	db *gorm.DB
}

func NewLocation(client sql.Client) ILocation {
	return &Location{db: client.GetConnection()}
}

func (r *Location) RegisterLocation(location entities.Location) error {
	return r.db.Save(&location).Error
}
