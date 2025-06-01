package models

import "gorm.io/gorm"

type Tax struct {
	gorm.Model
	TaxRate int `json:"taxRate"`
}
