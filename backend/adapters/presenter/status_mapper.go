package presenter

import (
	"radar/entities"
	"radar/presenters"
)

func MapCreateStatusEntityToPresenter(s entities.Status, l entities.Location) presenters.CreateStatus{

	return presenters.CreateStatus{
		Status:   presenters.Status{
			ID: s.ID,
			Name: s.Name,
			ProfileID: s.ProfileID,
			CreatedAt: s.CreatedAt,
		},
		Location: presenters.Location{
				ID:l.ID,
				ProfileID:l.ProfileID,
				Longitude:l.Longitude,
				Latitude:l.Latitude,
		},
	}
}