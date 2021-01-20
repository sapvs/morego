package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num > 2 && num%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Floor(math.Sqrt(float64(num)))); i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	count, _ := strconv.ParseInt(sc.Text(), 10, 32)

	var nums []int64
	for ; count > 0; count-- {
		sc.Scan()
		num, _ := strconv.ParseInt(sc.Text(), 10, 32)
		nums = append(nums, num)
	}

	for _, num := range nums {
		if isPrime(int(num)) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}

}
