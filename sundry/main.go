package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello")
	})

	log.Fatalf("Error in serving %v", http.ListenAndServe(":8080", nil))
}
