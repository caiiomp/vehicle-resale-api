package responses

import (
	"time"

	"github.com/caiomp87/vehicle-resale-api/src/core/domain/entity"
)

type Vehicle struct {
	ID        string    `json:"id"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	Year      int       `json:"year"`
	Color     string    `json:"color"`
	Price     float64   `json:"price"`
	Sold      bool      `json:"sold"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func VehicleFromDomain(vehicle entity.Vehicle) Vehicle {
	return Vehicle{
		ID:        vehicle.ID,
		Brand:     vehicle.Brand,
		Model:     vehicle.Model,
		Year:      vehicle.Year,
		Color:     vehicle.Color,
		Price:     vehicle.Price,
		Sold:      vehicle.Sold,
		CreatedAt: vehicle.CreatedAt,
		UpdatedAt: vehicle.UpdatedAt,
	}
}
