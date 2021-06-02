package bike

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func (ref *GormRepository) List() (result Bikes, err error) {
	err = ref.db.Find(&result).Error
	return
}

func (ref *GormRepository) Get(bikeID uuid.UUID) (result *Bike, err error) {
	tx := ref.db.First(&result, bikeID)

	if tx.Error == gorm.ErrRecordNotFound {
		err = ErrBikeNotFound
		return
	}

	err = tx.Error
	return
}

func (ref *GormRepository) Create(bike Bike) (result Bike, err error) {
	if bike.ID == uuid.Nil {
		bike.ID = uuid.New()
	}

	err = ref.db.Create(&bike).Error
	result = bike
	return
}

func (ref *GormRepository) Delete(bike Bike) (err error) {
	ref.db.Delete(&bike)
	return
}
