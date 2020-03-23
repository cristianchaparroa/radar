package initializer

import (
	"github.com/jinzhu/gorm"
	"radar/providers/sql"
	"radar/entities"
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
}
