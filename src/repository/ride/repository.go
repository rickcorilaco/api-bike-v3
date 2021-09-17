package ride

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	"github.com/rickcorilaco/api-bike-v3/src/core/values"
	"gorm.io/gorm"
)

var (
	ErrRideNotFound = errors.New("ride not found")
)

type Repository interface {
	List(filter values.RideListFilter) (result []entity.Ride, err error)
	Get(rideID uuid.UUID) (result *entity.Ride, err error)
	Create(ride entity.Ride) (result *entity.Ride, err error)
	Delete(ride entity.Ride) (err error)
}

func New(db interface{}) (gormRepository Repository, err error) {
	switch v := db.(type) {
	case *gorm.DB:
		gormRepository = &GormRepository{db: v}
	default:
		err = errors.New("invalid repository")
	}

	return
}
