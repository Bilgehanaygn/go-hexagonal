package response

import (
	"github.com/google/uuid"
)

type CategoryUpdateResponse struct {
	Id uuid.UUID `json:"id"`
}
