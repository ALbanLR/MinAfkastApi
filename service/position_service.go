package services

import (
	"errors"
	"udbytter/database"
	"udbytter/models"
)

func GetAllPositions() ([]models.Position, error) {
	var positions []models.Position
	if err := database.DB.Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}

func CreatePosition(position *models.Position) error {
	position.TotalValue = position.ShareNumber * position.SharePrice
	return database.DB.Create(position).Error
}

func DeletePosition(id string) error {
	var position models.Position
	if err := database.DB.First(&position, id).Error; err != nil {
		return errors.New("position not found")
	}
	return database.DB.Delete(&position).Error
}

func UpdatePosition(id string, input *models.Position) (*models.Position, error) {
	var position models.Position
	if err := database.DB.First(&position, id).Error; err != nil {
		return nil, errors.New("position not found")
	}

	position.Label = input.Label
	position.Description = input.Description
	position.ShareNumber = input.ShareNumber
	position.SharePrice = input.SharePrice
	position.TotalValue = input.ShareNumber * input.SharePrice
	position.Yield = input.Yield

	if err := database.DB.Save(&position).Error; err != nil {
		return nil, err
	}

	return &position, nil
}

func CalculateAverageYield() (float32, error) {
	var positions []models.Position
	if err := database.DB.Find(&positions).Error; err != nil {
		return 0, err
	}

	if len(positions) == 0 {
		return 0, nil
	}

	var totalYield float32
	for _, p := range positions {
		totalYield += p.Yield
	}

	return totalYield / float32(len(positions)), nil
}

func CalculateMonthlyIncome() (float32, error) {
	var positions []models.Position
	if err := database.DB.Find(&positions).Error; err != nil {
		return 0, err
	}

	var monthlyIncome float32
	for _, p := range positions {
		monthlyIncome += ((p.Yield / float32(12)) / 100) * float32(p.TotalValue)
	}

	return monthlyIncome, nil
}
