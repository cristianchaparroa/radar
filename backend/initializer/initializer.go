package initializer

import (
	"github.com/jinzhu/gorm"
	"radar/entities"
	"radar/providers/sql"
)

type Initializer struct {
	db *gorm.DB
}

func NewInitializer(client sql.Client) *Initializer {
	return &Initializer{db: client.GetConnection()}
}

func (i *Initializer) CreateTables() {
	i.db.CreateTable(&entities.Profile{})
	i.db.CreateTable(&entities.Location{})
	i.db.CreateTable(&entities.Status{})
}
