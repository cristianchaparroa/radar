package presenter

import (
	"radar/entities"
	"radar/presenters"
)

func NewProfileResponse(profile entities.Profile, status *entities.Status) presenters.ProfileResponse{

	return presenters.ProfileResponse{

		Profile: presenters.Profile{
			ID:profile.ID,
			DeviceID:profile.DeviceID,
			CreatedAt:profile.CreatedAt,
		},
		Status:  presenters.Status{
			ID : status.ID,
			ProfileID:profile.ID,
			Name: status.Name,
			CreatedAt:status.CreatedAt,
		},

	}
}

func MapProfilePresenterToEntity(profile presenters.Profile) entities.Profile {

	return entities.Profile{
		ID:        profile.ID,
		DeviceID:  profile.DeviceID,
		CreatedAt: profile.CreatedAt,
	}
}