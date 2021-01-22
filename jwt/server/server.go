package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/sudosapan/morego/jwt/common"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Authorized World")
}

func checkTokenAuth(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Token")
		if tokenString == "" {
			fmt.Fprintf(w, "Not Authorized : No Token")
			return
		}

		token, err := jwt.Parse(tokenString, common.GetKey)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		if token.Valid {
			endpoint(w, r)
		}
	})
}

func main() {
	http.Handle("/", checkTokenAuth(homePage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
