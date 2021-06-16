package ride

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	List() (result Rides, err error)
	Get(rideID uuid.UUID) (result *Ride, err error)
	Create(ride Ride) (result Ride, err error)
	Delete(ride Ride) (err error)
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
