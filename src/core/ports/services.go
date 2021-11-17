package ports

import (
	"github.com/google/uuid"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

type BikeService interface{
	List(filter value.BikeListFilter) (result *domain.Bikes, err error)
	Get(bikeID uuid.UUID) (result *domain.Bike, err error)
	Create(bike domain.Bike) (result *domain.Bike, err error)
	Delete(bike domain.Bike) (result *domain.Bike, err error)
}

type RideService interface{
	List(filter value.RideListFilter) (result *domain.Rides, err error)
	Get(rideID uuid.UUID) (result *domain.Ride, err error)
	Create(ride domain.Ride) (result *domain.Ride, err error)
	Delete(ride domain.Ride) (result *domain.Ride, err error)
}