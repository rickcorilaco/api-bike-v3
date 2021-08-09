package values

import "github.com/google/uuid"

type BikeListFilter struct {
	ID    *uuid.UUID `query:"id"`
	Brand *string    `query:"brand"`
	Model *string    `query:"model"`
}
