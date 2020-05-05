package model

import (
	"log"
	"reflect"
	"testing"

	mocket "github.com/selvatico/go-mocket"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func SetupTests() *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	db, err := gorm.Open(mocket.DriverName, "connection_string")
	if err != nil {
		log.Fatalf("error mocking gorm: %s", err)
	}
	return db
}

func getMockGorm(enableLogs bool) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gDB, _ := gorm.Open("mysql", db)
	return gDB, mock
}

func TestUser_getUser(t *testing.T) {
	mockDB, sqlMock := getMockGorm(true)

	userMockRowsTest := sqlMock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "username"}).AddRow("1", "2020-05-04 16:09:34", "2020-05-04 16:09:34", nil, "Test")
	userMockRowsNotTest := sqlMock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "username"}).AddRow("1", "2020-05-04 16:09:34", "2020-05-04 16:09:34", nil, "NotTest")

	/*
		+----+---------------------+---------------------+------------+----------+
		| id | created_at          | updated_at          | deleted_at | username |
		+----+---------------------+---------------------+------------+----------+
		|  1 | 2020-05-04 16:09:34 | 2020-05-04 16:09:34 | NULL       | Test     |
		|  2 | 2020-05-04 16:09:34 | 2020-05-04 16:09:34 | NULL       | NotTest  |
		+----+---------------------+---------------------+------------+----------+
	*/

	//sqlMock.ExpectBegin()
	sqlMock.ExpectQuery("SELECT (.*) FROM users  WHERE (.+) username = \\?)) (.*)").WithArgs("Test").WillReturnRows(userMockRowsTest)
	sqlMock.ExpectQuery("SELECT (.*) FROM users  WHERE (.+) username = \\?)) (.*)").WithArgs("NotTest").WillReturnRows(userMockRowsNotTest)

	//sqlMock.ExpectClose()

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		user *User
		args args
		want User
	}{
		{
			name: "test",
			user: &User{Username: "Test"},
			args: args{db: mockDB},
			want: User{Username: "Test"},
		},
		{
			name: "notTest",
			user: &User{Username: "NotTest"},
			args: args{db: mockDB},
			want: User{Username: "NotTest"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUser(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.getUser() = %v, want %v", got, tt.want)
			}
		})
	}
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_getUsers(t *testing.T) {
	mockDB := SetupTests()

	/*
		+----+---------------------+---------------------+------------+----------+
		| id | created_at          | updated_at          | deleted_at | username |
		+----+---------------------+---------------------+------------+----------+
		|  1 | 2020-05-04 16:09:34 | 2020-05-04 16:09:34 | NULL       | Test     |
		|  2 | 2020-05-04 16:09:34 | 2020-05-04 16:09:34 | NULL       | NotTest  |
		+----+---------------------+---------------------+------------+----------+
	*/
	commonReply := []map[string]interface{}{{"id": 1, "username": "Test"}, {"id": 2, "username": "NotTest"}}
	mocket.Catcher.Reset().NewMock().WithQuery("SELECT * FROM \"users\"  WHERE \"users\".\"deleted_at\" IS NULL").WithReply(commonReply)
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		user *User
		args args
		want []User
	}{
		{
			name: "test",
			user: &User{Username: "Test"},
			args: args{db: mockDB},
			want: []User{User{Username: "Test"},
				User{Username: "NotTest"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUsers(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.getUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
