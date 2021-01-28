package fizzbuzz

func fizzbuzzV1(num int) []interface{} {
	var ret []interface{}
	for i := 1; i <= num; i++ {
		if i%15 == 0 {
			ret = append(ret, "fizzbuzz")
		} else if i%5 == 0 {
			ret = append(ret, "buzz")
		} else if i%3 == 0 {
			ret = append(ret, "fizz")
		} else {
			ret = append(ret, i)
		}
	}

	return ret
}

func fizzbuzzV2(num int) []interface{} {
	var ret []interface{}
	for i := 1; i <= num; i++ {
		app := ""
		if i%3 == 0 {
			app += "fizz"
		}
		if i%5 == 0 {
			app += "buzz"
		}
		if app == "" {
			ret = append(ret, i)
		} else {
			ret = append(ret, app)
		}
	}
	return ret
}

func fizzbuzzV3(num int) []interface{} {
	var ret []interface{}
	num3, num5 := 0, 0
	for i := 1; i <= num; i++ {
		app := ""
		num3++
		num5++
		if num3 == 3 {
			app += "fizz"
			num3 = 0
		}
		if num5 == 5 {
			app += "buzz"
			num5 = 0
		}
		if app == "" {
			ret = append(ret, i)
		} else {
			ret = append(ret, app)
		}
	}
	return ret
}
