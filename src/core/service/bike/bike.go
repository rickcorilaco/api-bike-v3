package bike

import (
	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	bikeRepository "github.com/rickcorilaco/api-bike-v3/src/repository/bike"
)

type BikeService struct {
	repository bikeRepository.Repository
}

func New(repository bikeRepository.Repository) (bikeService *BikeService, err error) {
	bikeService = &BikeService{repository: repository}
	return
}

func (ref *BikeService) List() (result []entity.Bike, err error) {
	repositoryResult, err := ref.repository.List()
	if err != nil {
		return
	}

	result = repositoryResult.ToDomain()
	return
}

func (ref *BikeService) Get(bikeID uuid.UUID) (result entity.Bike, err error) {
	repositoryResult, err := ref.repository.Get(bikeID)
	if err != nil {
		return
	}

	result = repositoryResult.ToDomain()
	return
}

func (ref *BikeService) Create(bike entity.Bike) (result entity.Bike, err error) {
	repoBike := bikeRepository.Bike{}
	repoBike.FromDomain(bike)

	repositoryResult, err := ref.repository.Create(repoBike)
	if err != nil {
		return
	}

	result = repositoryResult.ToDomain()
	return
}

func (ref *BikeService) Delete(bike entity.Bike) (err error) {
	repoBike := bikeRepository.Bike{}
	repoBike.FromDomain(bike)

	err = ref.repository.Delete(repoBike)
	return
}
