package service

import (
	"fmt"

	bikeService "github.com/rickcorilaco/api-bike-v3/src/core/service/bike"
	bikeRepository "github.com/rickcorilaco/api-bike-v3/src/repository/bike"
)

type Config struct {
	Repositories []interface{}
}

func Start(config Config) (services []interface{}, err error) {
	for _, repository := range config.Repositories {
		switch value := repository.(type) {
		case bikeRepository.Repository:
			var service *bikeService.BikeService

			service, err = bikeService.New(value)
			if err != nil {
				return
			}

			services = append(services, service)

		default:
			err = fmt.Errorf("invalid repository: %v", repository)
			break
		}
	}

	return
}
