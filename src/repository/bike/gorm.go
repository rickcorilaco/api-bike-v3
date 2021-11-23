package bike

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

type GormBikeRepository struct {
	db *gorm.DB
}

func NewGormBikeRepository(db *gorm.DB) (gormBikeRepository *GormBikeRepository, err error) {
	gormBikeRepository = &GormBikeRepository{db: db}
	return
}

func (ref *GormBikeRepository) List(filter value.BikeListFilter) (result *domain.Bikes, err error) {
	bike := Bike{}
	bike.FromListFilter(filter)

	model := Bikes{}

	err = ref.db.Where(&bike).Find(&model).Error
	if err != nil {
		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormBikeRepository) Get(bikeID uuid.UUID) (result *domain.Bike, err error) {
	model := &Bike{}

	err = ref.db.First(&model, bikeID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}

		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormBikeRepository) Create(bike domain.Bike) (result *domain.Bike, err error) {
	if bike.ID == uuid.Nil {
		bike.ID = uuid.New()
	}

	model := &Bike{}
	model.FromDomain(&bike)

	err = ref.db.Create(model).Error
	if err != nil {
		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormBikeRepository) Delete(bike domain.Bike) (err error) {
	var model = Bike{}
	model.FromDomain(&bike)

	err = ref.db.Delete(&model).Error
	return
}
