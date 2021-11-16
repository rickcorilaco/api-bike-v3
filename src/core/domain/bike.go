package domain

import uuid "github.com/google/uuid"

type Bikes []Bike

type Bike struct {
	ID    uuid.UUID
	Brand string
	Model string
}
