package postgres

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/response"
	"github.com/bilgehanaygn/urun/internal/pkg/postgres"
)

func toDbEntity(category *domain.Category) *CategoryDbEntity {
	dbCategory := &CategoryDbEntity{
		BaseEntity: postgres.BaseEntity{
			Id: category.Id,
		},
		Name:   category.Name,
		Kind:   category.Kind,
		Status: category.Status,
	}

	return dbCategory
}

func toDomainEntity(dbCategory *CategoryDbEntity) *domain.Category {

	category := domain.Category{
		Id:               dbCategory.BaseEntity.Id,
		Name:             dbCategory.Name,
		Kind:             dbCategory.Kind,
		ParentCategoryId: dbCategory.ParentCategoryId,
		Status:           dbCategory.Status,
	}

	return &category
}

func toCategoryDetailDto(dbCategory *CategoryDbEntity) *response.CategoryDetailDto {
	dto := response.CategoryDetailDto{
		Id:               dbCategory.BaseEntity.Id,
		Name:             dbCategory.Name,
		Kind:             dbCategory.Kind,
		ParentCategoryId: dbCategory.ParentCategoryId,
		Status:           dbCategory.Status,
	}

	return &dto
}
