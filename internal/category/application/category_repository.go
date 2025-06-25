package application

import "github.com/bilgehanaygn/urun/internal/category/domain"

type CategoryRepository interface {
	Find(id string) (*domain.Category, error)
	Save(s *domain.Category) error
	Update(s *domain.Category) error
}