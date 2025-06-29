package request

import (
	"fmt"

	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
)

type CategoryCreateRequest struct {
	Name             string              `json:"name"`
	Kind             domain.CategoryKind `json:"kind"`
	ParentCategoryId *uuid.UUID          `json:"parentCategoryId,omitempty"`
}

func (request *CategoryCreateRequest) ToDomainEntity() (*domain.Category, error) {
	if request.Kind == domain.MAIN_CATEGORY {
		if request.ParentCategoryId != nil {
			return nil, fmt.Errorf("main category cannot have a parent category")
		}
	} else {
		if request.ParentCategoryId == nil && request.Kind != domain.MAIN_CATEGORY {
			return nil, fmt.Errorf("%q must have a parent category", request.Kind)
		}
	}

	return &domain.Category{
		Id:               uuid.New(),
		Name:             request.Name,
		Kind:             domain.CategoryKind(request.Kind),
		ParentCategoryId: request.ParentCategoryId,
		Status:           domain.ACTIVE,
	}, nil
}
