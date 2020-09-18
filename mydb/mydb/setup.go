package mydb

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sudosapan/db/model"
)

const (
	dbName string = "sapan"
	dbUser string = "sapan"
	dbPwd  string = "abc123"
)

var (
	db            *gorm.DB = nil
	dbInitialized uint32
	dbInitLock    sync.Mutex
)

func createRows(db *gorm.DB) {
	//dropCreateTables(db)
	order1 := model.Order{Price: 123}
	order2 := model.Order{Price: 1234}
	usr := model.User{Username: "Test", Order: []model.Order{order1, order2}}
	db.NewRecord(&usr)
	db.Create(&usr)

	order1 = model.Order{Price: 223}
	order2 = model.Order{Price: 2234}
	usr = model.User{Username: "NotTest", Order: []model.Order{order1, order2}}
	db.NewRecord(&usr)
	db.Create(&usr)

}

func dropCreateTables(db *gorm.DB) {
	if db.HasTable(&model.User{}) {
		db.DropTable(&model.User{})
	}
	db.CreateTable(&model.User{})

	if db.HasTable(&model.Order{}) {
		db.DropTable(&model.Order{})
	}
	db.CreateTable(&model.Order{})
}

func GetConnectionWithLog(enableLogs bool) (*gorm.DB, error) {

	fmt.Println(fmt.Sprintf("intiilaxed flag %v", dbInitialized))

	if atomic.LoadUint32(&dbInitialized) == 1 {
		fmt.Println(fmt.Sprintf("intiilaxed flag is 1 "))
		if err := db.DB().Ping(); err == nil {
			return db, nil
		} else {
			fmt.Println(fmt.Sprintf("Error in ping %v", err))
		}
		atomic.StoreUint32(&dbInitialized, 0)
	}

	fmt.Println(fmt.Sprintf("Checking if need to create database as it is %v and intiilaxed flag %v", db, dbInitialized))

	dbInitLock.Lock()
	defer dbInitLock.Unlock()

	if atomic.LoadUint32(&dbInitialized) == 0 {
		fmt.Println(fmt.Sprintf("Creating database as it is %v ", db))
		params := "?charset=utf8&parseTime=True&loc=Local"
		var conString = fmt.Sprintf("%s:%s@/%s%s", dbUser, dbPwd, dbName, params)
		fmt.Println(fmt.Sprintf("Connecting new to watcher db : %s ...\n", conString))

		var err error
		db, err = gorm.Open("mysql", conString)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error connecting to database  : %v ", err))
			return db, err
		}

		db.LogMode(enableLogs)
		db.DB().SetMaxOpenConns(50)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetConnMaxLifetime(time.Minute * 30)

		atomic.StoreUint32(&dbInitialized, 1)
		fmt.Println(fmt.Sprintf("Set initailized flag to %v", dbInitialized))
	}
	return db, nil
}
