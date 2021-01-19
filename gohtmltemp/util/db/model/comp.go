package model

import "gorm.io/gorm"

// Company company
type Company struct {
	gorm.Model
	Name string `json:"company_name" gorm:"name"`
}
