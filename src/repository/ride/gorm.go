package ride

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	"github.com/rickcorilaco/api-bike-v3/src/core/values"
)

type GormRepository struct {
	db *gorm.DB
}

func (ref *GormRepository) List(filter values.RideListFilter) (result []entity.Ride, err error) {
	query, args := ref.generateQueryAndArgsFromFilter(filter)
	model := Rides{}

	err = ref.db.Where(query, args...).Find(&model).Error
	if err != nil {
		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormRepository) Get(rideID uuid.UUID) (result *entity.Ride, err error) {
	model := &Ride{}
	tx := ref.db.First(&model, rideID)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			err = ErrRideNotFound
			return
		}

		err = tx.Error
		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormRepository) Create(ride entity.Ride) (result *entity.Ride, err error) {
	if ride.ID == uuid.Nil {
		ride.ID = uuid.New()
	}

	model := &Ride{}
	model.FromDomain(&ride)

	err = ref.db.Create(model).Error
	if err != nil {
		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormRepository) Delete(ride entity.Ride) (err error) {
	var model = Ride{}
	model.FromDomain(&ride)

	err = ref.db.Delete(&model).Error
	return
}

func (ref *GormRepository) generateQueryAndArgsFromFilter(filter values.RideListFilter) (query string, args []interface{}) {
	parts := []string{}
	args = []interface{}{}

	if filter.DistanceGreaterThan != nil {
		parts = append(parts, "distance >= ?")
		args = append(args, *filter.DistanceGreaterThan)
	}

	if filter.DistanceLessThan != nil {
		query += "distance <= ?"
		args = append(args, *filter.DistanceLessThan)
	}

	query = strings.Join(parts, " AND ")
	return
}
