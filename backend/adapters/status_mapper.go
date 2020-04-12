package adapters

import (
	"radar/domain"
	"radar/models"
	"radar/presenters"
)

func MapCreateStatusDomainToPresenter(s domain.Status, l domain.Location) presenters.CreateStatus{

	return presenters.CreateStatus{
		Status:   presenters.Status{
			ID: s.ID,
			Name: s.Name,
			ProfileID: s.ProfileID,
			CreatedAt: s.CreatedAt,
			Current:s.Current,
		},
		Location: presenters.Location{
				ID:l.ID,
				ProfileID:l.ProfileID,
				Longitude:l.Longitude,
				Latitude:l.Latitude,
		},
	}
}



func MapStatusEntityToDomain(s models.Status) domain.Status {
	return domain.Status{
		ID:         s.ID,
		ProfileID:  s.ProfileID,
		//LocationID: s.LocationID,
		Name:       s.Name,
		CreatedAt:  s.CreatedAt,
		Current:    s.Current,
	}
}

func MapStatusesEntityToDomain(statuses []models.Status) []domain.Status {

	domains := make([]domain.Status,len(statuses))

	for i, s := range statuses {
		status := MapStatusEntityToDomain(s)
		domains[i] = status
	}

	return domains
}