package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndlsServe(":8080", nil)
}

func router() {}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
