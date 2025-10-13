package responses

import (
	"time"

	"github.com/caiiomp/vehicle-resale-api/src/core/domain/entity"
)

type Sale struct {
	ID        string    `json:"id,omitempty"`
	VehicleID string    `json:"vehicle_id"`
	UserID    string    `json:"user_id"`
	Price     float64   `json:"price"`
	SoldAt    time.Time `json:"sold_at"`
}

func SaleFromDomain(sale entity.Sale) Sale {
	return Sale{
		ID:        sale.ID,
		VehicleID: sale.VehicleID,
		UserID:    sale.UserID,
		Price:     sale.Price,
		SoldAt:    sale.SoldAt,
	}
}
