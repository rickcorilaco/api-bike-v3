package bike

import (
	"errors"

	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
)

var (
	ErrBikeNotFound = errors.New("bike not found")
)

func New(db interface{}) (bikeRepository ports.BikeRepository, err error) {
	switch v := db.(type) {
	case *gorm.DB:
		bikeRepository, err = NewGormBikeRepository(v)
	default:
		err = errors.New("invalid database")
	}

	return
}
