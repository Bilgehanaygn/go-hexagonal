package postgres

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/bilgehanaygn/urun/internal/pkg/postgres"
	"github.com/google/uuid"
)

type CategoryDbEntity struct {
	postgres.BaseEntity
	Name             string
	Kind             domain.CategoryKind
	Status           domain.ActivenessStatus
	ParentCategoryId *uuid.UUID
}

func (CategoryDbEntity) TableName() string {
	return "category"
}
