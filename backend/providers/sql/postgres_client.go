package sql

import (
	"fmt"
	//_ "github.com/bmizerany/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

type PostgresClient struct {
	db *gorm.DB
}

func (c *PostgresClient) GetConnection() *gorm.DB {
	return c.db
}

func NewPostgresClient() (Client, error) {

	user := os.Getenv("USER_DB")
	pass := os.Getenv("PASSWORD_DB")
	dbName := os.Getenv("NAME_DB")
	host := os.Getenv("HOST_DB")

	dataSource := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, pass, host, dbName)
	fmt.Println(dataSource)

	db, err := gorm.Open("postgres", dataSource)

	if err != nil {
		return nil, err
	}

	return &PostgresClient{db: db}, nil
}
