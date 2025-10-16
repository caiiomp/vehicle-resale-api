package vehicleRepository

import (
	"context"
	"sort"
	"time"

	"github.com/caiiomp/vehicle-resale-api/src/core/domain/entity"
	"github.com/caiiomp/vehicle-resale-api/src/repository/model"
	"github.com/google/uuid"
)

type VehicleRepository interface {
	Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id string) (*entity.Vehicle, error)
	Search(ctx context.Context, isSold *bool) ([]entity.Vehicle, error)
	Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error)
}

type vehicleRepository struct {
	vehicles []model.Vehicle
}

func NewVehicleRepository() VehicleRepository {
	return &vehicleRepository{
		vehicles: []model.Vehicle{},
	}
}

func (ref *vehicleRepository) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	record := model.VehicleFromDomain(vehicle)

	record.ID = uuid.NewString()

	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	ref.vehicles = append(ref.vehicles, record)

	for _, vehicle := range ref.vehicles {
		if vehicle.ID == record.ID {
			return vehicle.ToDomain(), nil
		}
	}

	return nil, nil
}

func (ref *vehicleRepository) GetByID(ctx context.Context, id string) (*entity.Vehicle, error) {
	for _, vehicle := range ref.vehicles {
		if vehicle.ID == id {
			return vehicle.ToDomain(), nil
		}
	}

	return nil, nil
}

func (ref *vehicleRepository) Search(ctx context.Context, isSold *bool) ([]entity.Vehicle, error) {
	var (
		hasFilter              bool
		filterJustSoldVehicles bool
	)

	if isSold != nil {
		hasFilter = true
		filterJustSoldVehicles = *isSold
	}

	vehicles := make([]entity.Vehicle, 0)

	for _, vehicle := range ref.vehicles {
		if hasFilter {
			if filterJustSoldVehicles {
				if vehicle.SoldAt != nil {
					vehicles = append(vehicles, *vehicle.ToDomain())
					continue
				}
			} else {
				if vehicle.SoldAt == nil {
					vehicles = append(vehicles, *vehicle.ToDomain())
					continue
				}
			}
		}

		vehicles = append(vehicles, *vehicle.ToDomain())
	}

	sort.Slice(vehicles, func(i, j int) bool {
		return vehicles[i].Price < vehicles[j].Price
	})

	return vehicles, nil
}

func (ref *vehicleRepository) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	vehicleIndex := -1

	for i, vehicle := range ref.vehicles {
		if vehicle.ID == id {
			vehicleIndex = i
			break
		}
	}

	if vehicleIndex == -1 {
		return nil, nil
	}

	var hasUpdate bool

	if vehicle.Brand != "" && vehicle.Brand != ref.vehicles[vehicleIndex].Brand {
		ref.vehicles[vehicleIndex].Brand = vehicle.Brand
		hasUpdate = true
	}

	if vehicle.Model != "" && vehicle.Model != ref.vehicles[vehicleIndex].Model {
		ref.vehicles[vehicleIndex].Model = vehicle.Model
		hasUpdate = true
	}

	if vehicle.Year != 0 && vehicle.Year != ref.vehicles[vehicleIndex].Year {
		ref.vehicles[vehicleIndex].Year = vehicle.Year
		hasUpdate = true
	}

	if vehicle.Color != "" && vehicle.Color != ref.vehicles[vehicleIndex].Color {
		ref.vehicles[vehicleIndex].Color = vehicle.Color
		hasUpdate = true
	}

	if vehicle.Price != 0 && vehicle.Price != ref.vehicles[vehicleIndex].Price {
		ref.vehicles[vehicleIndex].Price = vehicle.Price
		hasUpdate = true
	}

	if vehicle.SoldAt != nil && vehicle.SoldAt != ref.vehicles[vehicleIndex].SoldAt {
		ref.vehicles[vehicleIndex].SoldAt = vehicle.SoldAt
		hasUpdate = true
	}

	if hasUpdate {
		ref.vehicles[vehicleIndex].UpdatedAt = time.Now()
	}

	return ref.vehicles[vehicleIndex].ToDomain(), nil
}
