package vehicle

import (
	"context"

	"github.com/caiomp87/vehicle-resale-api/src/core/domain/entity"
	"github.com/caiomp87/vehicle-resale-api/src/repository/vehicleRepository"
)

type VehicleService interface {
	Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id string) (*entity.Vehicle, error)
	Search(ctx context.Context, isSold *bool) ([]entity.Vehicle, error)
	Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error)
}

type vehicleService struct {
	vehicleRepository vehicleRepository.VehicleRepository
}

func NewVehicleService(vehicleRepository vehicleRepository.VehicleRepository) VehicleService {
	return &vehicleService{
		vehicleRepository: vehicleRepository,
	}
}

func (ref *vehicleService) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	return ref.vehicleRepository.Create(ctx, vehicle)
}

func (ref *vehicleService) GetByID(ctx context.Context, id string) (*entity.Vehicle, error) {
	return ref.vehicleRepository.GetByID(ctx, id)
}

func (ref *vehicleService) Search(ctx context.Context, isSold *bool) ([]entity.Vehicle, error) {
	return ref.vehicleRepository.Search(ctx, isSold)
}

func (ref *vehicleService) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	return ref.vehicleRepository.Update(ctx, id, vehicle)
}
