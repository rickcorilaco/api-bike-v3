package ride

import (
	"time"

	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
)

type Ride struct {
	ID       uuid.UUID `json:"id"`
	Distance float64   `json:"distance" validate:"required"`
	Duration string    `json:"duration" validate:"required"`
	Date     time.Time `json:"date" validate:"required"`
}

func (ref *Ride) FromDomain(dom domain.Ride) {
	ref.ID = dom.ID
	ref.Distance = dom.Distance
	ref.Duration = dom.Duration.String()
	ref.Date = dom.Date
}

func (ref *Ride) ToDomain() (dom domain.Ride, err error) {
	dom = domain.Ride{
		ID:       ref.ID,
		Distance: ref.Distance,
		Date:     ref.Date,
	}

	if ref.Duration != "" {
		dom.Duration, err = time.ParseDuration(ref.Duration)
		if err != nil {
			return
		}
	}

	return
}

type Rides []Ride

func (ref *Rides) FromDomain(dom *domain.Rides) {
	for _, domRide := range *dom {
		ride := Ride{}
		ride.FromDomain(domRide)
		*ref = append(*ref, ride)
	}

	return
}

func (ref *Rides) ToDomain() (dom []domain.Ride, err error) {
	dom = []domain.Ride{}

	for _, refRide := range *ref {
		var ride domain.Ride

		ride, err = refRide.ToDomain()
		if err != nil {
			return
		}

		dom = append(dom, ride)
	}

	return
}
