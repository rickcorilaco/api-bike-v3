package bike

import (
	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

type Bike struct {
	ID    uuid.UUID `gorm:"id"`
	Brand string    `gorm:"brand"`
	Model string    `gorm:"model"`
}

func (ref *Bike) FromDomain(dom *domain.Bike) {
	if dom == nil {
		return
	}

	ref.ID = dom.ID
	ref.Brand = dom.Brand
	ref.Model = dom.Model
}

func (ref *Bike) ToDomain() (dom *domain.Bike) {
	dom = &domain.Bike{
		ID:    ref.ID,
		Brand: ref.Brand,
		Model: ref.Model,
	}

	return
}

func (ref *Bike) FromListFilter(filter value.BikeListFilter) {
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

func (ref *Bikes) FromDomain(dom []domain.Bike) {
	for _, domBike := range dom {
		bike := Bike{}
		bike.FromDomain(&domBike)
		*ref = append(*ref, bike)
	}
}

func (ref *Bikes) ToDomain() (dom *domain.Bikes) {
	dom = &domain.Bikes{}

	for _, refBike := range *ref {
		*dom = append(*dom, *refBike.ToDomain())
	}

	return
}
