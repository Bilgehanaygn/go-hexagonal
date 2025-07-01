package response

import (
	"github.com/google/uuid"
)

type CatalogUpdateResponse struct {
	Id uuid.UUID `json:"id"`
} 