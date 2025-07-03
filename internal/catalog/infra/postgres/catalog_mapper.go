package postgres

import (
	"github.com/bilgehanaygn/urun/internal/catalog/domain"
	"github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"
	"github.com/bilgehanaygn/urun/internal/pkg/postgres"
)

func toDbEntity(cat *domain.Catalog) *CatalogDbEntity {
	cpDbs := make([]CatalogProductDbEntity, len(cat.CatalogProducts))

	for i, cp := range cat.CatalogProducts {
		cpDbs[i] = CatalogProductDbEntity{
			BaseEntity: postgres.BaseEntity{Id: cp.Id},
			CatalogId:  cat.Id,
			ProductId:  cp.ProductId,
			Price:      cp.Price,
		}
	}

	return &CatalogDbEntity{
		BaseEntity:       postgres.BaseEntity{Id: cat.Id},
		Name:             cat.Name,
		CatalogProducts:  cpDbs,
	}
}

func toDomainEntity(dbCat *CatalogDbEntity) *domain.Catalog {
	cps := make([]*domain.CatalogProduct, len(dbCat.CatalogProducts))

	for i, cp := range dbCat.CatalogProducts {
		cps[i] = &domain.CatalogProduct{
			Id:        cp.BaseEntity.Id,
			CatalogId: cp.CatalogId,
			ProductId: cp.ProductId,
			Price:     cp.Price,
		}
	}

	return &domain.Catalog{
		Id:              dbCat.BaseEntity.Id,
		Name:            dbCat.Name,
		CatalogProducts: cps,
	}
}

func toCatalogDetailDto(dbCat *CatalogDbEntity) *response.CatalogDetailDto {
	cpDtos := make([]response.CatalogProductDto, len(dbCat.CatalogProducts))

	for i, cp := range dbCat.CatalogProducts {
		cpDtos[i] = response.CatalogProductDto{
			Id:        cp.BaseEntity.Id,
			CatalogId: dbCat.BaseEntity.Id,
			ProductId: cp.ProductId,
			Price:     cp.Price,
		}
	}

	return &response.CatalogDetailDto{
		Id:              dbCat.BaseEntity.Id,
		Name:            dbCat.Name,
		CatalogProducts: cpDtos,
	}
}