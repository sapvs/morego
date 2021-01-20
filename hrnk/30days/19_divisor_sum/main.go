package sum

func divisorSumV1(number int) int {
	var sum int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	return sum
}

func divisorSumV2(number int) int {
	if number == 1 {
		return 1
	}

	var sum int
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	return sum + 1 + number
}
