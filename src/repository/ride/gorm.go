package ride

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

type GormRideRepository struct {
	db *gorm.DB
}

func NewGormRideRepository(db *gorm.DB) (gormRideRepository *GormRideRepository, err error) {
	gormRideRepository = &GormRideRepository{db: db}
	return
}

func (ref *GormRideRepository) List(filter value.RideListFilter) (result *domain.Rides, err error) {
	query, args := ref.generateQueryAndArgsFromFilter(filter)
	model := Rides{}

	err = ref.db.Where(query, args...).Find(&model).Error
	if err != nil {
		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormRideRepository) Get(rideID uuid.UUID) (result *domain.Ride, err error) {
	model := &Ride{}

	err = ref.db.First(&model, rideID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
			return
		}

		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormRideRepository) Create(ride domain.Ride) (result *domain.Ride, err error) {
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

func (ref *GormRideRepository) Delete(ride domain.Ride) (err error) {
	var model = Ride{}
	model.FromDomain(&ride)

	err = ref.db.Delete(&model).Error
	return
}

func (ref *GormRideRepository) generateQueryAndArgsFromFilter(filter value.RideListFilter) (query string, args []interface{}) {
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
