package request

import "github.com/google/uuid"

type CatalogGetRequest struct {
	Id uuid.UUID `param:"id"`
}