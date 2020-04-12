package services

import (
	"radar/domain"
	"radar/repositories"
)

type CheckInfections interface {

}

type LocationCheckInfections struct {
	infections repositories.Infections
}

func (i *LocationCheckInfections) GetInfections() []domain.Profile{

	positives, err  := i.infections.GetAllPositiveCases()

	if err != nil {

	}

	contacts, err  := i.infections.GetContacts(positives)

	if err != nil {

	}

	return contacts
}