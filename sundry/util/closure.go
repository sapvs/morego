package util

func closure() func() int {
	count := 10
	return func() int {
		count = count - 1
		return count
	}
}
