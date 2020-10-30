package dao

import "log"

// FailOnError function to fail
func FailOnError(e error, msg string) {
	if e != nil {
		log.Fatalf("%s  due to %v ", msg, e)
	}
}
