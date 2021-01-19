package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"html/template"

	"github.com/sudosapan/html/model"
)

var templates *template.Template

func init() {
	log.Printf("Time before %v ", time.Nanosecond)
	templates = template.Must(template.ParseGlob("templates/*.html"))
	log.Printf("Time after %v ", time.Nanosecond)
}

func main() {
	register(&model.Route{Path: "/", Method: http.MethodGet, Handler: index})
	register(&model.Route{Path: "/hello", Method: http.MethodGet, Handler: hello})
	register(&model.Route{Path: "/post", Method: http.MethodGet, Handler: post})
	register(&model.Route{Path: "/login", Method: http.MethodGet, Handler: login})
	model.StartServe(nil, ":8080")
}

func register(route *model.Route) {
	if err := model.RegisterHandler(nil, route); err != nil {
		log.Fatalf("Could not register router for %s with method %s due to %v", route.Path, route.Method, err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login")
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post")
}
