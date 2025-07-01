package postgres

import (
	"github.com/bilgehanaygn/urun/internal/catalog/domain"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"
	"github.com/bilgehanaygn/urun/internal/common/postgres"
)

func toDbEntity(catalog *domain.Catalog) *CatalogDbEntity {
	dbCatalog := &CatalogDbEntity{
		BaseEntity: postgres.BaseEntity{
			ID: catalog.Id,
		},
		Name:     catalog.Name,
		// Products: convert to []uuid.UUID if needed
	}
	return dbCatalog
}

func toDomainEntity(dbCatalog *CatalogDbEntity) *domain.Catalog {
	catalog := domain.Catalog{
		Id:       dbCatalog.BaseEntity.ID,
		Name:     dbCatalog.Name,
		// Products: convert from []uuid.UUID if needed
	}
	return &catalog
}

func toCatalogDetailDto(dbCatalog *CatalogDbEntity) *response.CatalogDetailDto {
	dto := response.CatalogDetailDto{
		Id:   dbCatalog.BaseEntity.ID,
		Name: dbCatalog.Name,
		// Products: convert from []uuid.UUID if needed
	}
	return &dto
} 