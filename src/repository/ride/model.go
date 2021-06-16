package ride

import (
	"time"

	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
)

type Ride struct {
	ID       uuid.UUID
	Distance float64
	Duration time.Duration
	Date     time.Time
}

func (ref *Ride) FromDomain(dom entity.Ride) {
	ref.ID = dom.ID
	ref.Distance = dom.Distance
	ref.Duration = dom.Duration
	ref.Date = dom.Date
}

func (ref *Ride) ToDomain() (dom *entity.Ride) {
	if ref == nil {
		return
	}

	dom = &entity.Ride{
		ID:       ref.ID,
		Distance: ref.Distance,
		Duration: ref.Duration,
		Date:     ref.Date,
	}

	return
}

type Rides []Ride

func (ref *Rides) FromDomain(dom []entity.Ride) {
	for _, domRide := range dom {
		ride := Ride{}
		ride.FromDomain(domRide)
		*ref = append(*ref, ride)
	}

	return
}

func (ref *Rides) ToDomain() (dom []entity.Ride) {
	dom = []entity.Ride{}

	for _, refRide := range *ref {
		dom = append(dom, *refRide.ToDomain())
	}

	return
}
