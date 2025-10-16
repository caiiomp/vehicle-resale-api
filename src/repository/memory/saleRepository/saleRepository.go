package saleRepository

import (
	"context"

	interfaces "github.com/caiiomp/vehicle-resale-api/src/core/_interfaces"
	"github.com/caiiomp/vehicle-resale-api/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-api/src/repository/model"
	"github.com/google/uuid"
)

type saleRepository struct {
	sales []model.Sale
}

func NewSaleRepository() interfaces.SaleRepository {
	return &saleRepository{
		sales: []model.Sale{},
	}
}

func (ref *saleRepository) Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error) {
	record := model.SaleFromDomain(sale)
	record.ID = uuid.NewString()

	ref.sales = append(ref.sales, record)

	for _, sale := range ref.sales {
		if sale.ID == record.ID {
			return sale.ToDomain(), nil
		}
	}

	return nil, nil
}

func (ref *saleRepository) Search(ctx context.Context) ([]entity.Sale, error) {
	sales := make([]entity.Sale, len(ref.sales))

	for i, sale := range ref.sales {
		sales[i] = *sale.ToDomain()
	}

	return sales, nil
}
