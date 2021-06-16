package service

import (
	bikeService "github.com/rickcorilaco/api-bike-v3/src/core/service/bike"
	rideService "github.com/rickcorilaco/api-bike-v3/src/core/service/ride"
	"github.com/rickcorilaco/api-bike-v3/src/repository"
)

type Config struct {
	Repositories repository.Repositories
}

type Services struct {
	Bike *bikeService.BikeService
	Ride *rideService.RideService
}

func Start(config Config) (services Services, err error) {
	if config.Repositories.Bike == nil {
		return
	}

	services.Bike, err = bikeService.New(config.Repositories.Bike)
	if err != nil {
		return
	}

	services.Ride, err = rideService.New(config.Repositories.Ride)
	return
}
