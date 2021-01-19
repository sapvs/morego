package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sudosapan/goactions/util"
)

func main() {
	log.Println("starting")
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", rootGet).Methods(http.MethodGet)
	server := http.Server{Addr: ":8080", Handler: myRouter}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to start http due to %v", err)
	} else {
		log.Printf("HTTP stopped by shutdown %v", err)
	}
}

func rootGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Git actions \n ")
	fmt.Fprintf(w, "sum of 2 + 2 is  %d \n ", util.Sum(2, 2))
}
