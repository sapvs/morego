package main

import "fmt"

func printNum(n int) {
	if n > 1 {
		printNum(n - 1)
	}
	fmt.Println(n)
}

func main() {
	printNum(50)
}
