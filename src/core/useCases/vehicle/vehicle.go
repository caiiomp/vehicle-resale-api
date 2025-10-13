package vehicle

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/caiiomp/vehicle-resale-api/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-api/src/repository/saleRepository"
	"github.com/caiiomp/vehicle-resale-api/src/repository/vehicleRepository"
)

type VehicleService interface {
	Create(ctx context.Context, vehicle entity.Vehicle, roleType string) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id string) (*entity.Vehicle, error)
	Search(ctx context.Context, isSold *bool) ([]entity.Vehicle, error)
	Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error)
	Sell(ctx context.Context, vehicleID, userID string) (*entity.Vehicle, error)
}

type vehicleService struct {
	vehicleRepository vehicleRepository.VehicleRepository
	saleRepository    saleRepository.SaleRepository
}

func NewVehicleService(vehicleRepository vehicleRepository.VehicleRepository, saleRepository saleRepository.SaleRepository) VehicleService {
	return &vehicleService{
		vehicleRepository: vehicleRepository,
		saleRepository:    saleRepository,
	}
}

func (ref *vehicleService) Create(ctx context.Context, vehicle entity.Vehicle, roleType string) (*entity.Vehicle, error) {
	if roleType != "ADMIN" {
		return nil, fmt.Errorf("role '%s' not allowed for this action", roleType)
	}

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

func (ref *vehicleService) Sell(ctx context.Context, vehicleID, userID string) (*entity.Vehicle, error) {
	vehicle, err := ref.vehicleRepository.GetByID(ctx, vehicleID)
	if err != nil {
		return nil, err
	}

	if vehicle == nil {
		return nil, errors.New("vehicle not found")
	}

	if vehicle.SoldAt != nil {
		return nil, errors.New("vehicle already sold")
	}

	soldTime := time.Now()

	sale := entity.Sale{
		VehicleID: vehicleID,
		UserID:    userID,
		Price:     vehicle.Price,
		SoldAt:    soldTime,
	}

	_, err = ref.saleRepository.Create(ctx, sale)
	if err != nil {
		return nil, err
	}

	soldVehicle := entity.Vehicle{
		SoldAt: &soldTime,
	}

	return ref.vehicleRepository.Update(ctx, vehicleID, soldVehicle)
}
