package models

import "gorm.io/gorm"

type TaxRate struct {
	gorm.Model
	Name    string  `json:"name"`
	TaxRate float32 `json:"taxRate"`
}
