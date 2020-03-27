package entities

import (
	"github.com/gin-gonic/gin"
	"radar/entities"
)

func GetLocation(c *gin.Context) (*entities.Location, error) {

	var location *entities.Location
	err := c.BindJSON(&location)

	if err != nil {
		return nil, err
	}
	return location, err
}