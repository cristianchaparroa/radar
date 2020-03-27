package presenter

import (
	"radar/entities"
	"radar/presenters"
)

func MapLocationPresenterToEntity(location presenters.Location) entities.Location {
	return entities.Location{
		ID:        location.ID,
		ProfileID: location.ProfileID,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		CreateAt:  location.CreateAt,
		UpdatedAt: location.UpdatedAt,
	}
}
