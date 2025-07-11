package response

import (
	"github.com/google/uuid"
)

type ProductDetailDto struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
	Status string   `json:"status"`
}
