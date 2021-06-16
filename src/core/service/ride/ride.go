package ride

import (
	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	rideRepository "github.com/rickcorilaco/api-bike-v3/src/repository/ride"
)

type RideService struct {
	repository rideRepository.Repository
}

func New(repository rideRepository.Repository) (rideService *RideService, err error) {
	rideService = &RideService{repository: repository}
	return
}

func (ref *RideService) List() (result []entity.Ride, err error) {
	repositoryResult, err := ref.repository.List()
	if err != nil {
		return
	}

	result = repositoryResult.ToDomain()
	return
}

func (ref *RideService) Get(rideID uuid.UUID) (result *entity.Ride, err error) {
	repositoryResult, err := ref.repository.Get(rideID)
	if err != nil {
		return
	}

	result = repositoryResult.ToDomain()
	return
}

func (ref *RideService) Create(ride entity.Ride) (result *entity.Ride, err error) {
	repoRide := rideRepository.Ride{}
	repoRide.FromDomain(ride)

	repositoryResult, err := ref.repository.Create(repoRide)
	if err != nil {
		return
	}

	result = repositoryResult.ToDomain()
	return
}

func (ref *RideService) Delete(ride entity.Ride) (err error) {
	repoRide := rideRepository.Ride{}
	repoRide.FromDomain(ride)

	err = ref.repository.Delete(repoRide)
	return
}
