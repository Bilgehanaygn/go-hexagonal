package command

import (
	"fmt"

	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/google/uuid"
)

type CategoryCreateRequest struct {
	Name string `json:"isim"`
	Kind domain.CategoryKind `json:"kategoriTuru"`
	ParentCategoryId uuid.UUID `json:"ebeveynKategoriId"` 
}

func (request *CategoryCreateRequest) ToDomainEntity() (*domain.Category, error){
    if request.Kind == domain.MAIN_CATEGORY {
        if request.ParentCategoryId != uuid.Nil {
            return nil, fmt.Errorf("main category cannot have a parent category")
        }
    } else {
        if request.ParentCategoryId == uuid.Nil {
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