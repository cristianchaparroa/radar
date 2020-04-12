package adapters

import (
	"radar/domain"
	"radar/models"
	"radar/presenters"
)

func NewProfileResponse(profile domain.Profile, status *domain.Status) presenters.ProfileResponse {

	return presenters.ProfileResponse{

		Profile: presenters.Profile{
			ID:        profile.ID,
			DeviceID:  profile.DeviceID,
			CreatedAt: profile.CreatedAt,
		},
		Status: presenters.Status{
			ID:        status.ID,
			ProfileID: profile.ID,
			Name:      status.Name,
			CreatedAt: status.CreatedAt,
		},
	}
}

func MapProfilePresenterToDomain(profile presenters.Profile) domain.Profile {

	return domain.Profile{
		ID:        profile.ID,
		DeviceID:  profile.DeviceID,
		CreatedAt: profile.CreatedAt,
	}
}

func MapProfileEntityToDomain(profile models.Profile)  domain.Profile{
	return domain.Profile{
		ID:        profile.ID,
		DeviceID:  profile.DeviceID,
		CreatedAt: profile.CreatedAt,
	}
}

func MapProfileDomainToEntity(profile domain.Profile) models.Profile {
	return models.Profile{
		ID:        profile.ID,
		DeviceID:  profile.DeviceID,
		CreatedAt: profile.CreatedAt,
	}
}