package repositories

import (
	"github.com/jinzhu/gorm"
	"radar/adapters"
	"radar/domain"
	"radar/models"
	"radar/providers/sql"
)

type Statuses struct {
	db *gorm.DB
}

func NewStatus(sql sql.Client) Status {
	return &Statuses{db: sql.GetConnection()}
}

func (r *Statuses) Create(location domain.Location, statusName string) (*domain.Status, error) {
	l := adapters.MapLocationDomainToEntity(location)
	s := models.NewStatus(l, statusName)
	err := r.db.Save(&s).Error
	status := adapters.MapStatusEntityToDomain(*s)
	return &status, err
}

func (r *Statuses) FindLastByUserID(profileID string) (domain.Status, error) {

	var status models.Status
	err := r.db.
		Where("profile_id = ?", profileID).
		Order("created_at DESC").
		First(&status).
		Error
	s := adapters.MapStatusEntityToDomain(status)
	return s, err
}

func (r *Statuses) FindAllByUserID(userID string) ([]domain.Status, error) {

	var occurrences []models.Status
	err := r.db.
		Where("profile_id = ?", userID).
		Order("created_at DESC").
		Find(&occurrences).
		Error

	statuses := adapters.MapStatusesEntityToDomain(occurrences)
	return statuses, err
}

func (r *Statuses) InactivateLast(profileID string) error {
	err := r.db.
		Update("current = ?", false).
		Where("profile_id = ?", profileID).
		Order("created_at desc").
		Limit(1).
		Error
	return err
}