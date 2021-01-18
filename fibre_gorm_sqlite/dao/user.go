package dao

import (
	"github.com/jinzhu/gorm"
)

//User denotes user
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

//AllUsers to retrieve all users from database as json
func AllUsers() *[]User {
	var users []User
	DB.Find(&users)
	return &users
}

//GetUsersByName to retrieve for given id
func GetUsersByName(name string) *[]User {
	var users []User
	DB.Where(&User{Name: name}).Find(&users)
	return &users
}

//GetUsersByEmail to retrieve for given id
func GetUsersByEmail(email string) *[]User {
	var users []User
	DB.Where(&User{Email: email}).Find(&users)
	return &users
}

//GetUsersByID to retrieve for given id
func GetUsersByID(id uint) *[]User {
	var users []User
	DB.Find(&users, id)
	return &users
}
