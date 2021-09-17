package values

import (
	"github.com/google/uuid"
)

type BikeListFilter struct {
	ID    *uuid.UUID `query:"id"`
	Brand *string    `query:"brand"`
	Model *string    `query:"model"`
}

type RideListFilter struct {
	ID                  *uuid.UUID `query:"id"`
	DistanceGreaterThan *float64   `query:"distance_greater_than"`
	DistanceLessThan    *float64   `query:"distance_less_than"`
}
