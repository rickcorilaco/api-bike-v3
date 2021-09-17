package ride

import (
	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	"github.com/rickcorilaco/api-bike-v3/src/core/values"
	rideRepository "github.com/rickcorilaco/api-bike-v3/src/repository/ride"
)

type RideService struct {
	repository rideRepository.Repository
}

func New(repository rideRepository.Repository) (rideService *RideService, err error) {
	rideService = &RideService{repository: repository}
	return
}

func (ref *RideService) List(filter values.RideListFilter) (result []entity.Ride, err error) {
	return ref.repository.List(filter)
}

func (ref *RideService) Get(rideID uuid.UUID) (result *entity.Ride, err error) {
	return ref.repository.Get(rideID)
}

func (ref *RideService) Create(ride entity.Ride) (result *entity.Ride, err error) {
	return ref.repository.Create(ride)
}

func (ref *RideService) Delete(ride entity.Ride) (err error) {
	return ref.repository.Delete(ride)
}
