package sale

import (
	"context"

	"github.com/caiiomp/vehicle-resale-api/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-api/src/repository/saleRepository"
)

type SaleService interface {
	Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error)
	Search(ctx context.Context) ([]entity.Sale, error)
}

type saleService struct {
	saleRepository saleRepository.SaleRepository
}

func NewSaleService(saleRepository saleRepository.SaleRepository) SaleService {
	return &saleService{
		saleRepository: saleRepository,
	}
}

func (ref *saleService) Create(ctx context.Context, sale entity.Sale) (*entity.Sale, error) {
	return ref.saleRepository.Create(ctx, sale)
}

func (ref *saleService) Search(ctx context.Context) ([]entity.Sale, error) {
	return ref.saleRepository.Search(ctx)
}
