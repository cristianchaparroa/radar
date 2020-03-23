package sql

import "github.com/jinzhu/gorm"

type Client interface {
	GetConnection() *gorm.DB
}
