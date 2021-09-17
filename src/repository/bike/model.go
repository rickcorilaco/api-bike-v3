package bike

import (
	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	"github.com/rickcorilaco/api-bike-v3/src/core/values"
)

type Bike struct {
	ID    uuid.UUID `gorm:"id"`
	Brand string    `gorm:"brand"`
	Model string    `gorm:"model"`
}

func (ref *Bike) FromDomain(dom *entity.Bike) {
	if dom == nil {
		return
	}

	ref.ID = dom.ID
	ref.Brand = dom.Brand
	ref.Model = dom.Model
}

func (ref *Bike) ToDomain() (dom *entity.Bike) {
	dom = &entity.Bike{
		ID:    ref.ID,
		Brand: ref.Brand,
		Model: ref.Model,
	}

	return
}

func (ref *Bike) FromListFilter(filter values.BikeListFilter) {
	if filter.ID != nil {
		ref.ID = *filter.ID
	}

	if filter.Brand != nil {
		ref.Brand = *filter.Brand
	}

	if filter.Model != nil {
		ref.Model = *filter.Model
	}
}

type Bikes []Bike

func (ref *Bikes) FromDomain(dom []entity.Bike) {
	for _, domBike := range dom {
		bike := Bike{}
		bike.FromDomain(&domBike)
		*ref = append(*ref, bike)
	}
}

func (ref *Bikes) ToDomain() (dom []entity.Bike) {
	dom = []entity.Bike{}

	for _, refBike := range *ref {
		dom = append(dom, *refBike.ToDomain())
	}

	return
}
