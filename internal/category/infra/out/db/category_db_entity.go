package db

import (
	"github.com/bilgehanaygn/urun/internal/category/domain"
	"github.com/bilgehanaygn/urun/internal/common/infra/out/db/entity"
	"github.com/google/uuid"
)

type CategoryDbEntity struct {
	entity.BaseEntity
	Name             string
	Kind             domain.CategoryKind
	Status           domain.ActivenessStatus
	ParentCategoryId uuid.UUID
}
