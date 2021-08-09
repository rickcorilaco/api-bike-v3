package bike

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	"github.com/rickcorilaco/api-bike-v3/src/core/values"
)

type GormRepository struct {
	db *gorm.DB
}

func (ref *GormRepository) List(filter values.BikeListFilter) (result []entity.Bike, err error) {
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

func (ref *GormRepository) Get(bikeID uuid.UUID) (result *entity.Bike, err error) {
	model := &Bike{}
	tx := ref.db.First(&model, bikeID)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			err = ErrBikeNotFound
			return
		}

		err = tx.Error
		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormRepository) Create(bike entity.Bike) (result *entity.Bike, err error) {
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

func (ref *GormRepository) Delete(bike entity.Bike) (err error) {
	var model *Bike
	model.FromDomain(&bike)

	err = ref.db.Delete(&model).Error
	return
}
