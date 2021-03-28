package model

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var defaultRouter *mux.Router

func init() {
	defaultRouter = mux.NewRouter()
}

//Route an http route
type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

//RegisterHandler registers route to router or to default is not provided
func RegisterHandler(router *mux.Router, route *Route) error {
	if router == nil {
		log.Printf("Using default route for %s %s \n", route.Path, route.Method)
		defaultRouter.HandleFunc(route.Path, route.Handler).Methods(route.Method)

	} else {
		if route.Path == "" {
			return fmt.Errorf("Route %v missing Path, Handler or Method", route)
		}
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
	return nil
}

//TODO
func validateRoute(route *Route) error {
	return nil
}

//StartServe start http listening, at give port
func StartServe(router *mux.Router, port string) {
	var thisRouter *mux.Router
	if router == nil {
		log.Printf("Using default router\n")
		thisRouter = defaultRouter
	} else {
		thisRouter = router
	}
	http.ListenAndServe(port, thisRouter)
}
