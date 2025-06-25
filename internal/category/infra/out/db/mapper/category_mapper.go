package mapper

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/bilgehanaygn/urun/internal/category/infra/out/db/entity"
	"github.com/bilgehanaygn/urun/internal/common/infra/out/db"
)

func toDbCategory(category *domain.Category) *entity.CategoryDbEntity {
	dbCategory := &entity.CategoryDbEntity{
		BaseEntity: db.BaseEntity{
            ID: category.Id,      
        },
        Name:   category.Name,
        Kind:   category.Kind,
        Status: category.Status,
	}

	return dbCategory
}

func fromDbCategory(dbCategory *entity.CategoryDbEntity) *domain.Category {
	category := domain.Category{
		Id: dbCategory.ID,
		Name: dbCategory.Name,
		Kind: dbCategory.Kind,
		ParentCategoryId: dbCategory.ParentCategoryId,
		Status: dbCategory.Status,
	}

	return &category
}