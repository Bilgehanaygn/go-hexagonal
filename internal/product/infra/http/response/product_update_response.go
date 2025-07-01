package response

import (
	"github.com/google/uuid"
)

type ProductUpdateResponse struct {
	Id uuid.UUID `json:"id"`
} 