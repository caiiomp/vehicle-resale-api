package entity

import "time"

type Sale struct {
	ID        string
	VehicleID string
	UserID    string
	Price     float64
	SoldAt    time.Time
}
