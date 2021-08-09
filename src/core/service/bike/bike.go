package bike

import (
	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	"github.com/rickcorilaco/api-bike-v3/src/core/values"
	bikeRepository "github.com/rickcorilaco/api-bike-v3/src/repository/bike"
)

type BikeService struct {
	repository bikeRepository.Repository
}

func New(repository bikeRepository.Repository) (bikeService *BikeService, err error) {
	bikeService = &BikeService{repository: repository}
	return
}

func (ref *BikeService) List(filter values.BikeListFilter) (result []entity.Bike, err error) {
	return ref.repository.List(filter)
}

func (ref *BikeService) Get(bikeID uuid.UUID) (result *entity.Bike, err error) {
	return ref.repository.Get(bikeID)
}

func (ref *BikeService) Create(bike entity.Bike) (result *entity.Bike, err error) {
	return ref.repository.Create(bike)
}

func (ref *BikeService) Delete(bike entity.Bike) (err error) {
	return ref.repository.Delete(bike)
}
