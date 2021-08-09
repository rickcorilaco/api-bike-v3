package bike

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	"github.com/rickcorilaco/api-bike-v3/src/core/values"
)

var (
	ErrBikeNotFound = errors.New("bike not found")
)

type Repository interface {
	List(filter values.BikeListFilter) (result []entity.Bike, err error)
	Get(bikeID uuid.UUID) (result *entity.Bike, err error)
	Create(bike entity.Bike) (result *entity.Bike, err error)
	Delete(bike entity.Bike) (err error)
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
