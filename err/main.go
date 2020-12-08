package main

import "fmt"

type httpError struct {
	Code    int
	Message string
}

func (m httpError) Error() string {
	return fmt.Sprintf("HTTP Error Code %d , Message %s", m.Code, m.Message)
}

func throwError() error {
	return httpError{Code: 202, Message: "Helloee"}
}

func main() {
	err := throwError()
	if err != nil {
		fmt.Println(err)
	}
}
