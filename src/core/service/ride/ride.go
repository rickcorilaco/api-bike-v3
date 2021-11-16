package ride

import (
	"github.com/google/uuid"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

type Service struct {
	repository ports.RideRepository
}

func New(repository ports.RideRepository) (rideService ports.RideService, err error) {
	rideService = &Service{repository: repository}
	return
}

func (ref *Service) List(filter value.RideListFilter) (result *domain.Rides, err error) {
	return ref.repository.List(filter)
}

func (ref *Service) Get(rideID uuid.UUID) (result *domain.Ride, err error) {
	return ref.repository.Get(rideID)
}

func (ref *Service) Create(ride domain.Ride) (result *domain.Ride, err error) {
	return ref.repository.Create(ride)
}

func (ref *Service) Delete(ride domain.Ride) (err error) {
	return ref.repository.Delete(ride)
}
