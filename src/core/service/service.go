package service

import (
	bikeService "github.com/rickcorilaco/api-bike-v3/src/core/service/bike"
	bikeRepository "github.com/rickcorilaco/api-bike-v3/src/repository/bike"
)

type Config struct {
	DB interface{}
}

func Start(config Config) (services []interface{}, err error) {
	if config.DB == nil {
		return
	}

	bikeRepository, err := bikeRepository.New(config.DB)
	if err != nil {
		return
	}

	bikeService, err := bikeService.New(bikeRepository)
	if err != nil {
		return
	}

	services = append(services, bikeService)
	return
}
