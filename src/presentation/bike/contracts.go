package bike

import (
	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
)

type Bike struct {
	ID    uuid.UUID `json:"id"`
	Brand string    `json:"brand"`
	Model string    `json:"model"`
}

func (ref *Bike) FromDomain(dom entity.Bike) {
	ref.ID = dom.ID
	ref.Brand = dom.Brand
	ref.Model = dom.Model
}

func (ref *Bike) ToDomain() (dom entity.Bike) {
	dom = entity.Bike{
		ID:    ref.ID,
		Brand: ref.Brand,
		Model: ref.Model,
	}

	return
}

type Bikes []Bike

func (ref *Bikes) FromDomain(dom []entity.Bike) {
	for _, domBike := range dom {
		bike := Bike{}
		bike.FromDomain(domBike)
		*ref = append(*ref, bike)
	}

	return
}

func (ref *Bikes) ToDomain() (dom []entity.Bike) {
	dom = []entity.Bike{}

	for _, refBike := range *ref {
		dom = append(dom, refBike.ToDomain())
	}

	return
}
