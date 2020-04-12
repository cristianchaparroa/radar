package adapters

import (
	"github.com/gin-gonic/gin"
	"radar/domain"
	"radar/models"
	"radar/presenters"
	"time"
)


func GetLocation(c *gin.Context) (*domain.Location, error) {

	var location *domain.Location
	err := c.BindJSON(&location)

	if err != nil {
		return nil, err
	}
	return location, err
}


func MapLocationPresenterToDomain(location presenters.Location) domain.Location {
	return domain.Location{
		ID:        location.ID,
		ProfileID: location.ProfileID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		CreateAt:  location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}
}

func MapLocationDomainToEntity(l domain.Location) models.Location {
	return models.Location{
		ID:        l.ID,
		ProfileID: l.ProfileID,
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
		CreateAt:  time.Time{},
		UpdatedAt: time.Time{},
	}
}