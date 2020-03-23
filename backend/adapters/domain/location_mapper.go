package domain

import (
	"github.com/gin-gonic/gin"
	"radar/domain"
	"radar/entities"
)

func GetLocation(c *gin.Context) (*domain.Location, error) {

	var location *domain.Location
	err := c.BindJSON(&location)

	if err != nil {
		return nil, err
	}

	return location, err
}

func MapDomainLocationToEntity(location domain.Location) *entities.Location {

	return &entities.Location{
		ID:        location.ID,
		UserID:    location.UserID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		CreateAt:  location.CreateAt,
		UpdatedAt: location.UpdatedAt,
	}
}


