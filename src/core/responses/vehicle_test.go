package responses

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-resale-api/src/core/domain/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestVehicleFromDomain(t *testing.T) {
	vehicleID := primitive.NewObjectID().Hex()

	now := time.Now()

	vehicle := entity.Vehicle{
		ID:        vehicleID,
		Brand:     "Some Brand",
		Model:     "Some Model",
		Year:      2025,
		Color:     "Gray",
		Price:     80000,
		SoldAt:    &now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	expected := Vehicle{
		ID:        vehicleID,
		Brand:     "Some Brand",
		Model:     "Some Model",
		Year:      2025,
		Color:     "Gray",
		Price:     80000,
		SoldAt:    &now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	actual := VehicleFromDomain(vehicle)

	assert.Equal(t, expected, actual)
}
