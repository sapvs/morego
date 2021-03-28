package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/live", live).Methods(http.MethodGet)
	router.HandleFunc("/ready", ready).Methods(http.MethodGet)
	log.Fatalf("Serving stopped %v", http.ListenAndServe(":8080", router))
}

func live(w http.ResponseWriter, r *http.Request) {
	log.Println("Live called ")
	fmt.Fprintln(w, "Live called")
}

func ready(w http.ResponseWriter, r *http.Request) {
	log.Println("Ready called ")
	fmt.Fprintln(w, "Ready called")
}
