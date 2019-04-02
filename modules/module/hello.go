package hello

import (
	"fmt"
	"rsc.io/quote/v3"
)

func Hello() string {
	fmt.Println(quote.HelloV3())
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
