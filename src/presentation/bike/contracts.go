package bike

import (
	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
)

type Bike struct {
	ID    uuid.UUID `json:"id"`
	Brand string    `json:"brand" validate:"required"`
	Model string    `json:"model" validate:"required"`
}

func (ref *Bike) FromDomain(dom domain.Bike) {
	ref.ID = dom.ID
	ref.Brand = dom.Brand
	ref.Model = dom.Model
}

func (ref *Bike) ToDomain() (dom domain.Bike) {
	dom = domain.Bike{
		ID:    ref.ID,
		Brand: ref.Brand,
		Model: ref.Model,
	}

	return
}

type Bikes []Bike

func (ref *Bikes) FromDomain(dom *domain.Bikes) {
	for _, domBike := range *dom {
		bike := Bike{}
		bike.FromDomain(domBike)
		*ref = append(*ref, bike)
	}

	return
}

func (ref *Bikes) ToDomain() (dom []domain.Bike) {
	dom = []domain.Bike{}

	for _, refBike := range *ref {
		dom = append(dom, refBike.ToDomain())
	}

	return
}
