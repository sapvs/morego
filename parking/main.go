package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Donee\n")
	case <-ctx.Done():
		log.Printf("Context Done %v", ctx.Err().Error())
		return
	}
}
