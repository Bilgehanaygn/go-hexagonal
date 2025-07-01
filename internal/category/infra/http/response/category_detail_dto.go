package response

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
)

type CategoryDetailDto struct {
	Id               uuid.UUID               `json:"id"`
	Name             string                  `json:"name"`
	Kind             domain.CategoryKind     `json:"kind"`
	ParentCategoryId *uuid.UUID              `json:"parentCategoryId"`
	Status           domain.ActivenessStatus `json:"status"`
}