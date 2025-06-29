package response

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
)

type CategoryDetailDto struct {
	id               uuid.UUID
	Name             string
	Kind             domain.CategoryKind
	ParentCategoryId uuid.UUID
	Status           domain.ActivenessStatus
}

func NewCategoryDetailDTO(category *domain.Category) CategoryDetailDto {
	return CategoryDetailDto{
		id:               category.Id,
		Name:             category.Name,
		Kind:             category.Kind,
		ParentCategoryId: category.ParentCategoryId,
		Status:           category.Status,
	}
}