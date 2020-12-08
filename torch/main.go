package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^(.+)@gmail$")
	match := re.FindAllStringSubmatch(r.URL.Path[1:], -1)
	msg := "hello, stranger"
	if match != nil {
		msg = "hello, " + match[0][1]
	}
	if _, err := fmt.Fprintln(w, msg); err != nil {
		log.Printf("could not write message: %v", err)
	}
}
