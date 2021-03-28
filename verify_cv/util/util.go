package util

import "log"

func LogFatal(err error, msg string) {
	if err != nil {
		log.Fatalf("%s due to %s ", msg, err.Error())
	}
}
