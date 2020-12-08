package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)

	t, _ := sc.ReadString('\n')

	num, e := strconv.Atoi(strings.TrimSuffix(t, "\n"))
	if e != nil {
		log.Fatalf("%v", e)
	}

	fmt.Printf("NUmber parsed %d\n", num)

	a, b := make([]float64, num), make([]float64, num)

	instr, e := sc.('\n')

	if e != nil {
		log.Fatalf("string 1 %v", e)
	}

	fmt.Printf("read string 1 %s\n", instr)

	for idx, s := range strings.Fields(strings.TrimSuffix(instr, "\n")) {
		a[idx], e = strconv.ParseFloat(s, 64)
		if e != nil {
			log.Fatalf("array in 1 %v", e)
		}
	}

	instr, e = sc.ReadString('\n')

	if e != nil {
		log.Fatalf("strin 2 %v", e)
	}

	fmt.Printf("read string 2 %s\n", instr)

	for idx, s := range strings.Fields(strings.TrimSuffix(instr, "\n")) {
		b[idx], e = strconv.ParseFloat(s, 64)
		if e != nil {
			log.Fatalf("arra in 2 %v", e)
		}
	}

	fmt.Printf("%.1f\n", WtdMean(a, b))

}

// WtdMean sxd
func WtdMean(a []float64, b []float64) float64 {
	if len(a) != len(b) {
		panic("Lenght not same")
	}

	var sumA, sumB float64
	for i := 0; i < len(a); i++ {
		sumA += (a[i] * b[i])
		sumB += b[i]
	}

	return sumA / sumB

}
