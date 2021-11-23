package ports

import (
	"github.com/google/uuid"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

type BikeRepository interface {
	List(filter value.BikeListFilter) (result *domain.Bikes, err error)
	Get(bikeID uuid.UUID) (result *domain.Bike, err error)
	Create(bike domain.Bike) (result *domain.Bike, err error)
	Delete(bike domain.Bike) (err error)
}

type RideRepository interface {
	List(filter value.RideListFilter) (result *domain.Rides, err error)
	Get(rideID uuid.UUID) (result *domain.Ride, err error)
	Create(ride domain.Ride) (result *domain.Ride, err error)
	Delete(ride domain.Ride) (err error)
}

type UserRepository interface {
	GetByUsername(username string) (result *domain.User, err error)
	Create(user domain.User) (result *domain.User, err error)
	ErrDuplicateRecord() (err error)
}
