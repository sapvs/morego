package common

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

var keys = map[string][]byte{
	"key": []byte("secretkeyforpropertykey"),
}

//GetKey gets key for the token
func GetKey(token *jwt.Token) (interface{}, error) {
	log.Printf("Recevied request to getKey %v", token)

	return keys[token.Header["key"].(string)], nil
}
