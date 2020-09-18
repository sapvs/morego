package main

import (
	"fmt"

	"github.com/sudosapan/db/model"
	"github.com/sudosapan/db/mydb"
)

func main() {
	fmt.Println("Hello")
	db, err := mydb.GetConnectionWithLog(true)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user := model.User{Username: "Test"}

	//db.Where("username like (?)", "%est%").Preload("Order", "price like (?)", "123%").Find(&usrs)

	fmt.Println(user.GetUser(db))

	usrs := user.GetUsers(db)
	for _, user := range usrs {
		fmt.Print("User name " + user.Username)
		if len(user.Order) >= 1 {
			for _, order := range user.Order {
				fmt.Println(fmt.Sprintf("Price %f", order.Price))
			}
		} else {
			fmt.Println("  Not match Order")
		}
	}

}
