package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	customtype()
}

func db(name string) *gorm.DB {
	db, err := gorm.Open("sqlite3", name)
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}
	return db
}
func mysql() {
	type Country struct {
		gorm.Model
		Name          string
		ContinentName string
	}
	type Address struct {
		gorm.Model
		CounID  int
		Country Country `gorm:"foreignKey:CounID; references:ID"`
	}

	type User struct {
		gorm.Model
		Name    string
		AddID   int
		Address Address `gorm:"foreignKey:AddID;references:ID"`
	}

	db := db("mysql.db")
	defer db.Close()

	db.DropTableIfExists(
		&User{},
		&Address{},
		&Country{},
	)

	log.Println("Dropped")
	db.AutoMigrate(
		&User{},
		&Address{},
		&Country{},
	)

	log.Println("Migrated")

	user := User{Name: "Hello", Address: Address{Country: Country{Name: "India", ContinentName: "Asia"}}}
	log.Printf("user %v", user)

	db.Create(&user)

	log.Println("Created")

	var users []User
	db.Find(&users)

	log.Printf("users %v", users)
}

type ClientType uint8

const (
	USER ClientType = iota
	SERVER
)

type Client struct {
	ClientID   uint64 `gorm:"primaryKey"`
	UserID     uint8
	ClientType ClientType `gorm:"type:number"`
	CreatedAt  time.Time
}

func (t ClientType) String() string {
	fmt.Printf("String() called with %d\n", t)
	return [...]string{"User", "Service"}[t]
}

// func (t ClientType) Value() (driver.Value, error) {

// 	log.Printf("Value called with %d", t)
// 	return int64(2), nil

// }
// func (t *ClientType) Scan(v interface{}) error {
// 		*t = 2
// 	return nil
// }

func customtype() {

	db := db("customtype.db")
	defer db.Close()
	db.LogMode(false)

	db.DropTableIfExists(
		&Client{},
	)

	log.Println("Dropped")
	db.AutoMigrate(
		&Client{},
	)

	var client Client

	db.LogMode(true)
	fmt.Println("Creating")
	if err := db.Create(&Client{ClientID: 9, UserID: 8, ClientType: USER, CreatedAt: time.Now()}).Error; err != nil {
		log.Fatalf("Erro creating %v", err)
	}
	fmt.Println("Created")

	fmt.Println("Finding")

	if err := db.First(&client).Error; err != nil {
		log.Fatalf("Erro getting %v", err)
	}
	fmt.Printf("Found %v\n", client)

}
