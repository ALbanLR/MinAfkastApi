package services

import (
	"udbytter/database"
	"udbytter/models"
)

func GetAllTaxRate() ([]models.TaxRate, error) {
	var taxRate []models.TaxRate
	if err := database.DB.Find(&taxRate).Error; err != nil {
		return nil, err
	}
	return taxRate, nil
}
