package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubso(arr []int64) int {
	var swaps int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				//swap
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				swaps++
			}
		}
	}
	return swaps
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()

	num, _ := strconv.ParseInt(scn.Text(), 10, 32)

	scn.Scan()
	spl := strings.Split(scn.Text(), " ")

	arr := make([]int64, num)

	for i := 0; i < int(num); i++ {
		arr[i], _ = strconv.ParseInt(spl[i], 10, 32)
	}

	fmt.Printf("Array is sorted in %d swaps.\n", bubso(arr))
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[num-1])
}
