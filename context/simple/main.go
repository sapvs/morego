package main

import (
	"fmt"
	"time"
)

func main() {

	printAfter(5*time.Second, "msg string")

	

}

func printAfter(t time.Duration, msg string) {
	time.Sleep(t)
	fmt.Println(msg)
}
