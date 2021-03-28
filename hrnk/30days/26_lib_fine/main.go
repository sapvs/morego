package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type date struct {
	day   int
	month int
	year  int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	returDate := readDate(sc)
	dueDate := readDate(sc)

	if returDate.year == dueDate.year {
		if returDate.month == dueDate.month {
			if returDate.day == dueDate.day {
				fmt.Printf("%d\n", 0)
			} else {
				if returDate.day > dueDate.day {
					fmt.Printf("%d\n", 15*(returDate.day-dueDate.day))
				} else {
					fmt.Printf("%d\n", 0)
				}
			}

		} else {
			if returDate.month > dueDate.month {
				fmt.Printf("%d\n", 500*(returDate.month-dueDate.month))
			} else {
				fmt.Printf("%d\n", 0)
			}
		}
	} else {
		if returDate.year > dueDate.year {
			fmt.Printf("%d\n", 10000)
		} else {
			fmt.Printf("%d\n", 0)
		}
	}

}

func getDiff(past *date, future *date) *date {
	return &date{
		day:   int(math.Abs(float64(future.day) - float64(past.day))),
		month: int(math.Abs(float64(future.month) - float64(past.month))),
		year:  int(math.Abs(float64(future.year) - float64(past.year))),
	}
}

func readDate(sc *bufio.Scanner) *date {
	sc.Scan()
	split := strings.Split(sc.Text(), " ")

	return &date{day: toInt(split[0]), month: toInt(split[1]), year: toInt(split[2])}
}

func toInt(in string) int {
	val, _ := strconv.ParseInt(in, 10, 32)
	return int(val)
}
