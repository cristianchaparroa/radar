package repositories

import (
	"radar/domain"
)

type Location interface {
	Register(location domain.Location) error
}

