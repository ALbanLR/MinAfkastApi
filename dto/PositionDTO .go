package dto

import (
	"udbytter/models"

	"gorm.io/gorm"
)

type PositionDTO struct {
	gorm.Model
	Label       string     `json:"label"`
	Description string     `json:"description"`
	ShareNumber float32    `json:"shareNumber"`
	SharePrice  float32    `json:"sharePrice"`
	TotalValue  float32    `json:"totalValue"`
	Yield       float32    `json:"yield"`
	TaxRate     TaxRateDTO `json:"taxRate"`
}

func ToPositionDTO(p models.Position, t models.TaxRate) PositionDTO {
	return PositionDTO{
		Model:       p.Model,
		Label:       p.Label,
		Description: p.Description,
		ShareNumber: p.ShareNumber,
		SharePrice:  p.SharePrice,
		TotalValue:  p.TotalValue,
		Yield:       p.Yield,
		TaxRate:     ToTaxRateDTO(t),
	}
}
