package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sudosapan/sqlite_gorm/dao"
	"github.com/sudosapan/sqlite_gorm/model"
)

func main() {
	fmt.Println("Starting web server")
	dao.Initialize()
	handleRequests()
}
func allUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	dao.DB.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")

	email := getVar(r, "email")
	name := getVar(r, "name")

	dao.DB.Create(&model.User{Email: email, Name: name})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
	var user model.User
	name := getVar(r, "name")

	dao.DB.Where(&model.User{Name: name}).Find(&user)

	fmt.Println("{}", user)

	dao.DB.Delete(&user)
	fmt.Fprintf(w, "User Delted")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}

func getVar(r *http.Request, name string) string {
	vars := mux.Vars(r)
	return vars[name]
}

func handleRequests() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/users", allUsers).Methods("GET")
	rtr.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	rtr.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	rtr.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", rtr))
}
