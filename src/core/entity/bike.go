package entity

import uuid "github.com/google/uuid"

type Bike struct {
	ID    uuid.UUID
	Brand string
	Model string
}
