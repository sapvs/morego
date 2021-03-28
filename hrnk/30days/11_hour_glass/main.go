package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func hour(arr [][]int) {
	numcol, numrow := len(arr[0]), len(arr)
	sum := math.MinInt64

	//For rows till 2 less
	for i := 0; i < numrow-2; i++ {
		for j := 0; j < numcol-2; j++ {
			sumIJ, _ := hourglass(arr, i, j)
			if sumIJ > sum {
				sum = sumIJ
			}
		}
	}

	fmt.Printf("%d", sum)
}

func hourglass(arr [][]int, i int, j int) (int, [][]int) {
	ret := [][]int{{arr[i][j], arr[i][j+1], arr[i][j+2]}, {arr[i+1][j+1]}, {arr[i+2][j], arr[i+2][j+1], arr[i+2][j+2]}}
	sum := int(arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2])

	return sum, ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := arrItemTemp
			arrRow = append(arrRow, int(arrItem))
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	hour(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
