package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

func main() {

	sc := bufio.NewReader(os.Stdin)

	t, _ := sc.ReadString('\n')

	num, e := strconv.Atoi(strings.TrimSuffix(t, "\n"))
	if e != nil {
		log.Fatalf("%v", e)
	}

	nums := make([]float64, num)
	instr, e := sc.ReadString('\n')

	if e != nil {
		log.Fatalf("%v", e)
	}

	for idx, s := range strings.Fields(strings.TrimSuffix(instr, "\n")) {
		nums[idx], e = strconv.ParseFloat(s, 64)
		if e != nil {
			log.Fatalf("%v", e)
		}
	}

	mean, e := stats.Mean(nums)
	if e != nil {
		log.Fatalf("%v", e)
	}
	median, e := stats.Median(nums)
	if e != nil {
		log.Fatalf("%v", e)
	}
	mode, e := stats.Mode(nums)
	if e != nil {
		log.Fatalf("%v", e)
	}
	var mode1 float64
	if len(mode) == 0 {
		sort.Float64s(nums)
		mode1 = nums[0]
	} else {
		mode1 = mode[0]
	}

	fmt.Printf("%.1f\n%.1f\n%.0f\n", mean, median, mode1)

}
