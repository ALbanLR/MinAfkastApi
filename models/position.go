package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Label       string  `json:"label"`
	Description string  `json:"description"`
	ShareNumber float32 `json:"shareNumber"`
	SharePrice  float32 `json:"sharePrice"`
	TotalValue  float32 `json:"totalValue"`
	Yield       float32 `json:"yield"`
	TaxRateID   uint    `json:"taxRateId"`
}
