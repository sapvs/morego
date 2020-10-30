package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Started")
	log.Fatal(startHTTP())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello simple http")
}

func startHTTP() *error {
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		return &err
	}
	return nil

}
