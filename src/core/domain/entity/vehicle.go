package entity

import "time"

type Vehicle struct {
	ID        string
	Brand     string
	Model     string
	Year      int
	Color     string
	Price     float64
	Sold      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
