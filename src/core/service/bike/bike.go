package bike

import (
	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/ports"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

type Service struct {
	repository ports.BikeRepository
}

func New(repository ports.BikeRepository) (bikeService ports.BikeService, err error) {
	bikeService = &Service{repository: repository}
	return
}

func (ref *Service) List(filter value.BikeListFilter) (result *domain.Bikes, err error) {
	return ref.repository.List(filter)
}

func (ref *Service) Get(bikeID uuid.UUID) (result *domain.Bike, err error) {
	return ref.repository.Get(bikeID)
}

func (ref *Service) Create(bike domain.Bike) (result *domain.Bike, err error) {
	return ref.repository.Create(bike)
}

func (ref *Service) Delete(bike domain.Bike) (result *domain.Bike, err error) {
	result, err = ref.Get(bike.ID)
	if err != nil {
		return
	}

	if result == nil {
		return
	}

	err = ref.repository.Delete(bike)
	return
}
