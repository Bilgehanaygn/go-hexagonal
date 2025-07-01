package response

import (
	"github.com/google/uuid"
)

type ProductCreateResponse struct {
	Id uuid.UUID `json:"id"`
} 