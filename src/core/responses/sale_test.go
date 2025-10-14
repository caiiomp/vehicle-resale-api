package responses

import (
	"testing"
	"time"

	"github.com/caiiomp/vehicle-resale-api/src/core/domain/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSaleFromDomain(t *testing.T) {
	saleID := primitive.NewObjectID().Hex()
	vehicleID := primitive.NewObjectID().Hex()
	userID := primitive.NewObjectID().Hex()

	now := time.Now()

	sale := entity.Sale{
		ID:        saleID,
		VehicleID: vehicleID,
		UserID:    userID,
		Price:     80000,
		SoldAt:    now,
	}

	expected := Sale{
		ID:        saleID,
		VehicleID: vehicleID,
		UserID:    userID,
		Price:     80000,
		SoldAt:    now,
	}

	actual := SaleFromDomain(sale)

	assert.Equal(t, expected, actual)
}
