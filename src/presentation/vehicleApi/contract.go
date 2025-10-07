package vehicleApi

import (
	"time"

	"github.com/caiomp87/vehicle-resale-api/src/core/domain/entity"
)

type createVehicleRequest struct {
	Brand string  `json:"brand" binding:"required"`
	Model string  `json:"model" binding:"required"`
	Year  int     `json:"year" binding:"required"`
	Color string  `json:"color" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func (ref createVehicleRequest) ToDomain() *entity.Vehicle {
	return &entity.Vehicle{
		Brand: ref.Brand,
		Model: ref.Model,
		Year:  ref.Year,
		Color: ref.Color,
		Price: ref.Price,
	}
}

type vehicleResponse struct {
	ID        string    `json:"id"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	Year      int       `json:"year"`
	Color     string    `json:"color"`
	Price     float64   `json:"price"`
	IsSold    bool      `json:"is_sold"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func vehicleResponseFromDomain(vehicle entity.Vehicle) vehicleResponse {
	return vehicleResponse{
		ID:        vehicle.ID,
		Brand:     vehicle.Brand,
		Model:     vehicle.Model,
		Year:      vehicle.Year,
		Color:     vehicle.Color,
		Price:     vehicle.Price,
		IsSold:    *vehicle.IsSold,
		CreatedAt: vehicle.CreatedAt,
		UpdatedAt: vehicle.UpdatedAt,
	}
}

type vehicleURI struct {
	ID string `uri:"vehicle_id"`
}

type updateVehicleRequest struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Year  int     `json:"year"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

func (ref updateVehicleRequest) ToDomain() *entity.Vehicle {
	return &entity.Vehicle{
		Brand: ref.Brand,
		Model: ref.Model,
		Year:  ref.Year,
		Color: ref.Color,
		Price: ref.Price,
	}
}

type vehicleQuery struct {
	IsSold *bool `form:"is_sold"`
}
