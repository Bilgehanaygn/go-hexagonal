package request

import (
	"github.com/google/uuid"
)

type ProductGetRequest struct {
	Id uuid.UUID `param:"id"`
}
