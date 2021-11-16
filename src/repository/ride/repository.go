package ride

import (
	"errors"

	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
)

var (
	ErrRideNotFound = errors.New("ride not found")
)

func New(db interface{}) (rideRepository ports.RideRepository, err error) {
	switch v := db.(type) {
	case *gorm.DB:
		rideRepository, err = NewGormRideRepository(v)
	default:
		err = errors.New("invalid database")
	}

	return
}
