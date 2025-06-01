package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Label       string  `json:"label"`
	Description string  `json:"description"`
	ShareNumber int     `json:"shareNumber"`
	SharePrice  int     `json:"sharePrice"`
	TotalValue  int     `json:"totalValue"`
	Yield       float32 `json:"yield"`
	//TaxRate     Tax     `json:"TaxRate,omitempty"`
}
