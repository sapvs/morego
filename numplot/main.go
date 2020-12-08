package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/sudosapan/numplot/myplot"
)

const (
	imageFile string = "plot.png"
	dataLen   int    = 999
)

func main() {

	data := make([]float64, dataLen)
	for i := 0; i < dataLen; i++ {
		data[i] = rand.Float64() * float64(100.0)
	}

	log.Printf("Plotting to file %s", imageFile)
	f, err := os.Create(imageFile)
	if err != nil {
		log.Fatalf(("could not open file %s due to %v"), imageFile, err)
	}

	defer f.Close()

	if err := myplot.Plot(f, data); err != nil {
		log.Fatalf("Could not plot %v", err)
	}

}
