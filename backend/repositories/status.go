package repositories

import (
	"github.com/jinzhu/gorm"
	"radar/entities"
	"radar/providers/sql"
)

type IStatus interface {
	Create(userID, status string) (*entities.Status, error)
}

type Status struct {
	db *gorm.DB
}

func NewStatus(sql sql.Client) IStatus {
	return &Status{db: sql.GetConnection()}
}

func (r *Status) Create(userID, status string) (*entities.Status, error) {
	s := entities.NewStatus(userID, status)
	err := r.db.Save(&s).Error
	return s, err
}

func (r *Status) FindLastByUserID(userID string) (entities.Status, error) {

	var status entities.Status
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		First(&status).
		Error
	return status, err
}

func (r *Status) FindAllByUserID(userID string) ([]entities.Status, error) {

	var statuses []entities.Status
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&statuses).
		Error
	return statuses, err
}
