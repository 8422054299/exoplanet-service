package models

type Exoplanet struct {
	ID          uint    `gorm:"primary_key"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Distance    int     `json:"distance" validate:"required,gt=10,lt=1000"`
	Radius      float64 `json:"radius" validate:"required,gt=0.1,lt=10"`
	Mass        float64 `json:"mass,omitempty"`
	Type        string  `json:"type" validate:"required,oneof=GasGiant Terrestrial"`
}
