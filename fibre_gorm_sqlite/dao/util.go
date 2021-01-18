package dao

import "log"

// Fail function to fail
func Fail(e error, msg string) {
	if e != nil {
		log.Fatalf("%s  due to %v ", msg, e)
	}
}

// Error function to log error
func Error(e error, msg string) {
	if e != nil {
		log.Printf("%s  due to %v \n", msg, e)
	}
}
