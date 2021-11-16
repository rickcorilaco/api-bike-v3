package domain

import (
	"time"

	uuid "github.com/google/uuid"
)

type Rides []Ride

type Ride struct {
	ID       uuid.UUID
	Distance float64
	Duration time.Duration
	Date     time.Time
}
