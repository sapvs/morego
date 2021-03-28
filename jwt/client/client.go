package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sudosapan/morego/jwt/common"
)

var (
	client *http.Client
	server string = os.Getenv("SERVER_URL")
)

func main() {
	client = http.DefaultClient
	http.HandleFunc("/", clientRoot)
	log.Fatalf("Http server stopped %v", http.ListenAndServe(":8090", nil))
}

func clientRoot(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodGet, server, nil)
	if err != nil {
		log.Printf("Request to %s failed due to %v", server, err)
	}
	req.Header.Set("Token", createJWT())
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Request to %s failed due to %v", server, err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	fmt.Fprintf(w, string(body))
}

func createJWT() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["some"] = "other"
	claims["exp"] = time.Now().Add(time.Second * 720).Unix()

	token.Header["key"] = "key"

	signingKey, err := common.GetKey(token)

	stringToken, err := token.SignedString(signingKey)
	if err != nil {
		panic(err)
	}

	return stringToken
}
