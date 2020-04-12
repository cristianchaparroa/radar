package initializer

import (
	"github.com/jinzhu/gorm"
	"radar/domain"
	"radar/providers/sql"
)

type Initializer struct {
	db *gorm.DB
}

func NewInitializer(client sql.Client) *Initializer {
	return &Initializer{db: client.GetConnection()}
}

func (i *Initializer) CreateTables() {
	i.db.CreateTable(&domain.Profile{})
	i.db.CreateTable(&domain.Location{})
	i.db.CreateTable(&domain.Status{})

	i.db.Model(&domain.Location{}).AddIndex("idx_geolocation", "latitude", "longitude")
}
