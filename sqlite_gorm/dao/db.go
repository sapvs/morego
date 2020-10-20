package dao

import (
	"log"

	"github.com/sudosapan/sqlite_gorm/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
		panic("failed to open connection to database")
	}
	defer db.Close()
	db.AutoMigrate(&model.User{})
}

func getUsers() {
}
