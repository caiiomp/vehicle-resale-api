package vehicleRepository

import (
	"context"

	"github.com/caiomp87/vehicle-resale-api/src/core/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type VehicleRepository interface {
	Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id string) (*entity.Vehicle, error)
	Search(ctx context.Context) ([]entity.Vehicle, error)
	Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error)
}

type vehicleRepository struct {
	db *mongo.Database
}

func NewVehicleRepository(db *mongo.Database) VehicleRepository {
	return &vehicleRepository{
		db: db,
	}
}

func (ref *vehicleRepository) Create(ctx context.Context, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	panic("unimplemented")
}

func (ref *vehicleRepository) GetByID(ctx context.Context, id string) (*entity.Vehicle, error) {
	panic("unimplemented")
}

func (ref *vehicleRepository) Search(ctx context.Context) ([]entity.Vehicle, error) {
	panic("unimplemented")
}

func (ref *vehicleRepository) Update(ctx context.Context, id string, vehicle entity.Vehicle) (*entity.Vehicle, error) {
	panic("unimplemented")
}
