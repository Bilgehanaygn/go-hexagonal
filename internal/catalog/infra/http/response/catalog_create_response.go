package response

import (
	"github.com/google/uuid"
)

type CatalogCreateResponse struct {
	Id uuid.UUID `json:"id"`
}
