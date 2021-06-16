package ride

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func (ref *GormRepository) List() (result Rides, err error) {
	err = ref.db.Find(&result).Error
	return
}

func (ref *GormRepository) Get(rideID uuid.UUID) (result *Ride, err error) {
	tx := ref.db.First(&result, rideID)

	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	err = tx.Error
	return
}

func (ref *GormRepository) Create(ride Ride) (result Ride, err error) {
	if ride.ID == uuid.Nil {
		ride.ID = uuid.New()
	}

	err = ref.db.Create(&ride).Error
	result = ride
	return
}

func (ref *GormRepository) Delete(ride Ride) (err error) {
	ref.db.Delete(&ride)
	return
}
