package model

import (
	"gorm.io/gorm"
)

//Employee employee belonging to Company
type Employee struct {
	gorm.Model
	Name    string `json:"name" gorm:"name"`
	Company *Company
	Manager *Employee
	Team    []Employee
}
