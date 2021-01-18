package dao

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB connection to database
var DB *gorm.DB
var once sync.Once

func connectDB() {
	log.Println("Connecting sqlite database with test.db here.")
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	Fail(err, "Could not open database")
}

func init() {
	once.Do(connectDB)
	initialMigration()
}

// InitialMigration To migrate database structure
func initialMigration() {
	DB.AutoMigrate(&User{})
}
