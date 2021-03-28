package main

import "fmt"

func solution(inp []string) []string {
	var odd, even string
	var out []string
	for _, word := range inp {
		for i, ch := range word {
			if i%2 == 0 {
				even = even + string(ch)
			} else {
				odd = odd + string(ch)
			}
		}
		out = append(out, even+" "+odd)
	}
	return out
}

func main() {
	var T int
	var word string
	var in []string

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&word)
		in = append(in, word)
		fmt.Println(in)
		fmt.Println()
	}
	for _, out := range solution(in) {
		fmt.Println(out)
	}
}
