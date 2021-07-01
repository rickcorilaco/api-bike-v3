package service

import (
	"errors"

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
	if err = validateConfig(config); err != nil {
		return
	}

	if services.Bike, err = bikeService.New(config.Repositories.Bike); err != nil {
		return
	}

	services.Ride, err = rideService.New(config.Repositories.Ride)
	return
}

func validateConfig(config Config) (err error) {
	if config.Repositories.Bike == nil {
		return errors.New("bike repository is nil")
	}

	if config.Repositories.Ride == nil {
		return errors.New("ride repository is nil")
	}

	return
}
