package postgres

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	pkgDomain "github.com/bilgehanaygn/urun/internal/pkg/domain"
	"github.com/bilgehanaygn/urun/internal/pkg/postgres"

	"github.com/google/uuid"
)

type CategoryDbEntity struct {
	postgres.BaseEntity
	Name             string
	Kind             domain.CategoryKind
	Status           pkgDomain.ActivenessStatus
	ParentCategoryId *uuid.UUID
}

func (CategoryDbEntity) TableName() string {
	return "category"
}
