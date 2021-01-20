package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	mmap := make(map[string]string, num)

	for i := num; i > 0; i-- {
		scanner.Scan()
		splitkeyval := strings.Split(scanner.Text(), " ")
		mmap[splitkeyval[0]] = splitkeyval[1]
	}

	for scanner.Scan() {
		query := scanner.Text()
		if query == "" {
			break
		}
		if _, ok := mmap[query]; ok {
			fmt.Printf("%s=%s\n", query, mmap[query])
		} else {
			fmt.Println("Not found")
		}

	}
}
