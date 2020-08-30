package mydb

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func getMockGorm(enableLogs bool) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gDB, _ := gorm.Open("mysql", db)
	return gDB, mock
}

func TestGetConnectionWithLog(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPing()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectClose()

	gDB, gerr := gorm.Open("mysql", db)

	t.Logf("%v  %v %v", gDB, gerr, gDB.DB().Stats().MaxOpenConnections)

	type args struct {
		enableLogs bool
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{
			name: "enabled",
			args: args{enableLogs: true},
		},
		{
			name: "disabled",
			args: args{enableLogs: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConnectionWithLog(tt.args.enableLogs)
			defer got.Close()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConnectionWithLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("GetConnectionWithLog() = %v, want %v", got, tt.want)
			}
		})
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_dropCreateTables(t *testing.T) {

	gdb, mock := getMockGorm(true)
	mock.ExpectQuery("SHOW TABLES FROM `` WHERE `Tables_in_` = ?").WithArgs("users")
	mock.ExpectClose()
	mock.ExpectBegin()

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "enabled",
			args: args{db: gdb},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dropCreateTables(tt.args.db)
		})
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_createRows(t *testing.T) {
	gdb, mock := getMockGorm(true)
	defer gdb.Close()
	mock.ExpectBegin()
	mock.ExpectClose()

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Create",
			args: args{db: gdb},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createRows(tt.args.db)
		})
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
