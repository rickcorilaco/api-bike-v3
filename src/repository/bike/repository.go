package bike

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrBikeNotFound = errors.New("bike not found")
)

type Repository interface {
	List() (result Bikes, err error)
	Get(bikeID uuid.UUID) (result *Bike, err error)
	Create(bike Bike) (result Bike, err error)
	Delete(bike Bike) (err error)
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
