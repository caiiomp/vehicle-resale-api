package model

import (
	"time"

	"github.com/caiiomp/vehicle-resale-api/src/core/domain/entity"
)

type Sale struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	VehicleID string    `json:"vehicle_id" bson:"vehicle_id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Price     float64   `json:"price" bson:"price"`
	SoldAt    time.Time `json:"sold_at" bson:"sold_at"`
}

func SaleFromDomain(sale entity.Sale) Sale {
	return Sale{
		VehicleID: sale.VehicleID,
		UserID:    sale.UserID,
		Price:     sale.Price,
		SoldAt:    sale.SoldAt,
	}
}

func (ref *Sale) ToDomain() *entity.Sale {
	return &entity.Sale{
		ID:        ref.ID,
		VehicleID: ref.VehicleID,
		UserID:    ref.UserID,
		Price:     ref.Price,
		SoldAt:    ref.SoldAt,
	}
}
