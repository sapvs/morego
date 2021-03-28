package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Am the server or client start as it is.
	role := os.Getenv("ROLE")
	log.Printf("I am %s\n", role)
	if role == "SERVER" {
		router := mux.NewRouter()
		router.HandleFunc("/", rootHandler).Methods(http.MethodGet)
		log.Fatalf("%v", http.ListenAndServe(":8080", router))
	} else {
		// am client
		router := mux.NewRouter()
		router.HandleFunc("/", clientRootHandler).Methods(http.MethodGet)
		log.Fatalf("%v", http.ListenAndServe(":8081", router))
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("received request")
	fmt.Fprint(w, "You landed right")
}

func clientRootHandler(w http.ResponseWriter, r *http.Request) {
	var client http.Client = *http.DefaultClient
	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("Error in client %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can;t read body %v", err)
	}

	log.Printf("Body read %s", string(body))
}
