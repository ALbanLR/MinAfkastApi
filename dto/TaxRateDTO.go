package dto

import (
	"udbytter/models"

	"gorm.io/gorm"
)

type TaxRateDTO struct {
	gorm.Model
	Name    string  `json:"name"`
	TaxRate float32 `json:"taxRate"`
}

func ToTaxRateDTO(t models.TaxRate) TaxRateDTO {
	return TaxRateDTO{
		Model:   t.Model,
		Name:    t.Name,
		TaxRate: t.TaxRate,
	}
}
