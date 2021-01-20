package main

import "math"

// IsPrime returns true if num is prime number
func IsPrime(num int) bool {
	return isPrimev3(num)
}

// isPrimev1 worst
func isPrimev1(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// isPrimev2 better 1
func isPrimev2(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// isPrimev3 more better
func isPrimev3(num int) bool {
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
