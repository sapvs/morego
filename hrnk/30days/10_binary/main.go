package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	binary, _ := strconv.ParseInt(scan.Text(), 10, 32)

	max := 0
	count := 0

	for _, c := range strconv.FormatInt(binary, 2) {
		if c == '1' {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}

	fmt.Printf("%d\n", max)
}
