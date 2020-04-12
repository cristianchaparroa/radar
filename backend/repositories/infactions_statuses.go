package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"radar/domain"
	"radar/providers/sql"
)

const(

)

type InfectionsStatuses struct {
	db *gorm.DB
}

func NewInfectionsStatuses(sql sql.Client) Infections {
	return &InfectionsStatuses{db :sql.GetConnection()}
}

func (i *InfectionsStatuses) GetAllPositiveCases() ([]domain.CurrentProfile, error) {
	var cases []domain.CurrentProfile

	err := i.db.Model(&domain.Status{}).
		Where("current = ?", true).
		Where("name = ?", domain.PositiveStatus).Find(&cases).
		Error

	return cases, err
}

func  (i *InfectionsStatuses)  GetContacts(positives []domain.CurrentProfile) ([]domain.Profile, error) {

	contacts := make([]domain.Profile, 0)

	for _,profile := range positives {
		cs := i.GetContact(profile.Profile.ID)
		contacts = append( contacts, cs...)
	}

	return contacts, nil
}


func (i *InfectionsStatuses) GetContact(profileID string) []domain.Profile {

	contacts := make([]domain.Profile, 0)

	query := fmt.Sprintf(AroundLocationQuery, )

	i.db.Exec(query)

	return contacts
}
