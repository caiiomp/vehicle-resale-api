package model

import (
	"time"

	"github.com/caiomp87/vehicle-resale-api/src/core/domain/entity"
)

type Vehicle struct {
	ID        string    `json:"id,omitempty" bson:"id,omitempty"`
	Brand     string    `json:"brand,omitempty" bson:"brand,omitempty"`
	Model     string    `json:"model,omitempty" bson:"model,omitempty"`
	Year      int       `json:"year,omitempty" bson:"year,omitempty"`
	Color     string    `json:"color,omitempty" bson:"color,omitempty"`
	Price     float64   `json:"price,omitempty" bson:"price,omitempty"`
	Sold      bool      `json:"sold,omitempty" bson:"sold,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

func VehicleFromDomain(vehicle entity.Vehicle) Vehicle {
	return Vehicle{
		ID:    vehicle.ID,
		Brand: vehicle.Brand,
		Model: vehicle.Model,
		Year:  vehicle.Year,
		Color: vehicle.Color,
		Price: vehicle.Price,
		Sold:  vehicle.Sold,
	}
}

func (ref Vehicle) ToDomain() *entity.Vehicle {
	return &entity.Vehicle{
		ID:        ref.ID,
		Brand:     ref.Brand,
		Model:     ref.Model,
		Year:      ref.Year,
		Color:     ref.Color,
		Price:     ref.Price,
		Sold:      ref.Sold,
		CreatedAt: ref.CreatedAt,
		UpdatedAt: ref.UpdatedAt,
	}
}
