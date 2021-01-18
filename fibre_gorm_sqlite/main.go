package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sudosapan/morego/fibre_gorm_sqlite/dao"
)

func main() {
	log.Println("Starting web server")
	handleRequests()
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	dao.Error(json.NewEncoder(w).Encode(dao.AllUsers()), "Failed to get all users from database")
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")

	email := getVar(r, "email")
	name := getVar(r, "name")

	dao.DB.Create(&dao.User{Email: email, Name: name})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
	var user dao.User
	name := getVar(r, "name")

	dao.DB.Where(&dao.User{Name: name}).Find(&user)

	fmt.Println("{}", user)

	dao.DB.Delete(&user)
	fmt.Fprintf(w, "User Deleted")
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

	dao.Fail(http.ListenAndServe(":8081", rtr), "Server quit due to error")
}
