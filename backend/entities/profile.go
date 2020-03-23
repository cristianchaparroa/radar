package entities

import "time"

type Profile struct {
	ID       string `gorm:"primary_key" json:"id"`
	Email    string `db:"email"`
	Password string `db:"password" json:"-"`
	CreateAt time.Time
}

func NewProfile(id string) *Profile {
	return &Profile{ID: id, CreateAt: time.Now()}
}
