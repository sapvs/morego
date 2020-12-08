package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type School struct {
	gorm.Model
	Name    string
	Student []Student
}

type Student struct {
	gorm.Model
	Name   string
	School School
}

type User struct {
	gorm.Model
	Username string
	Order    []Order
}
type Order struct {
	gorm.Model
	UserID uint
	Price  float64
}

func (user *User) GetUser(db *gorm.DB) User {
	var retUser User
	//db.Where(&User{Username: user.Username}).Preload("Order", "price like (?)", "123%").First(&retUser)
	db.Where(&User{Username: user.Username}).First(&retUser)
	return retUser
}

func (user *User) GetUsers(db *gorm.DB) []User {
	var retUsers []User
	db.Find(&retUsers)
	return retUsers
}

func (user User) String() string {
	return fmt.Sprintf("Name: %s ", user.Username)
}
