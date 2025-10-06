package vehicle

import (
	"context"

	"github.com/caiomp87/vehicle-resale-api/src/core/domain/entity"
	"github.com/caiomp87/vehicle-resale-api/src/repository/vehicleRepository"
)

type VehicleService interface {
	Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id string) (*entity.Vehicle, error)
	Search(ctx context.Context) ([]entity.Vehicle, error)
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

func (v *vehicleService) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	panic("unimplemented")
}

func (v *vehicleService) GetByID(ctx context.Context, id string) (*entity.Vehicle, error) {
	panic("unimplemented")
}

func (v *vehicleService) Search(ctx context.Context) ([]entity.Vehicle, error) {
	panic("unimplemented")
}

func (v *vehicleService) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	panic("unimplemented")
}
