package service

import (
	"errors"

	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
	bikeService "github.com/rickcorilaco/api-bike-v3/src/core/service/bike"
	rideService "github.com/rickcorilaco/api-bike-v3/src/core/service/ride"
	userService "github.com/rickcorilaco/api-bike-v3/src/core/service/user"
	"github.com/rickcorilaco/api-bike-v3/src/repository"
)

type Config struct {
	Repositories repository.Repositories
	UserTokenKey string
}

type Services struct {
	Bike ports.BikeService
	Ride ports.RideService
	User ports.UserService
}

func Start(config Config) (services Services, err error) {
	if err = validateConfig(config); err != nil {
		return
	}

	if services.Bike, err = bikeService.New(config.Repositories.Bike); err != nil {
		return
	}

	if services.Ride, err = rideService.New(config.Repositories.Ride); err != nil {
		return
	}

	services.User, err = userService.New(config.Repositories.User, config.UserTokenKey)
	return
}

func validateConfig(config Config) (err error) {
	if config.Repositories.Bike == nil {
		return errors.New("bike repository is nil")
	}

	if config.Repositories.Ride == nil {
		return errors.New("ride repository is nil")
	}

	if config.Repositories.User == nil {
		return errors.New("user repository is nil")
	}

	if config.UserTokenKey == ""{
		return errors.New("user token key is empty")
	}

	return
}
