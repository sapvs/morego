package dao

import (
	"sync"

	"github.com/sudosapan/sqlite_gorm/model"

	"github.com/jinzhu/gorm"
)

// DB connection to database
var DB *gorm.DB
var once sync.Once

func connectDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	FailOnError(err, "Could not open database")
}

// Initialize To migrate database structure
func Initialize() {
	once.Do(connectDB)
	initialMigration()
}

// InitialMigration To migrate database structure
func initialMigration() {
	DB.AutoMigrate(&model.User{})
}
