package db

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/bilgehanaygn/urun/internal/common/db"
)

func toDbEntity(category *domain.Category) *CategoryDbEntity {
	dbCategory := &CategoryDbEntity{
		BaseEntity: db.BaseEntity{
			ID: category.Id,
		},
		Name:   category.Name,
		Kind:   category.Kind,
		Status: category.Status,
	}

	return dbCategory
}

func toDomainEntity(dbCategory *CategoryDbEntity) *domain.Category {

	category := domain.Category{
		Id:               dbCategory.BaseEntity.ID,
		Name:             dbCategory.Name,
		Kind:             dbCategory.Kind,
		ParentCategoryId: dbCategory.ParentCategoryId,
		Status:           dbCategory.Status,
	}

	return &category
}
