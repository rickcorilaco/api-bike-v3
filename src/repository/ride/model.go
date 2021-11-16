package ride

import (
	"time"

	"github.com/google/uuid"
	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
)

type Ride struct {
	ID       uuid.UUID     `gorm:"id"`
	Distance float64       `gorm:"distance"`
	Duration time.Duration `gorm:"duration"`
	Date     time.Time     `gorm:"date"`
}

func (ref *Ride) FromDomain(dom *domain.Ride) {
	if dom == nil {
		return
	}

	ref.ID = dom.ID
	ref.Distance = dom.Distance
	ref.Duration = dom.Duration
	ref.Date = dom.Date
}

func (ref *Ride) ToDomain() (dom *domain.Ride) {
	if ref == nil {
		return
	}

	dom = &domain.Ride{
		ID:       ref.ID,
		Distance: ref.Distance,
		Duration: ref.Duration,
		Date:     ref.Date,
	}

	return
}

type Rides []Ride

func (ref *Rides) FromDomain(dom []domain.Ride) {
	for _, domRide := range dom {
		ride := Ride{}
		ride.FromDomain(&domRide)
		*ref = append(*ref, ride)
	}
}

func (ref *Rides) ToDomain() (dom *domain.Rides) {
	dom = &domain.Rides{}

	for _, refRide := range *ref {
		*dom = append(*dom, *refRide.ToDomain())
	}

	return
}
