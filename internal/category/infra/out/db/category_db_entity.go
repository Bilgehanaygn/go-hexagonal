package db

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/bilgehanaygn/urun/internal/common/db"
	"github.com/google/uuid"
)

type CategoryDbEntity struct {
	db.BaseEntity
	Name             string
	Kind             domain.CategoryKind
	Status           domain.ActivenessStatus
	ParentCategoryId uuid.UUID
}
