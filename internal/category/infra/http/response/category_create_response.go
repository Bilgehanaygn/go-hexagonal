package response

import (
	"github.com/google/uuid"
)

type CategoryCreateResponse struct {
	Id uuid.UUID `json:"id"`
}
