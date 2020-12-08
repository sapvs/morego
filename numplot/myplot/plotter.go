package myplot

import (
	"image/color"
	"io"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

// Plot plots the data
func Plot(w io.Writer, data []float64) error {
	plot, err := plot.New()
	if err != nil {
		return err
	}

	//plot.Title.TextStyle = draw.TextStyle{Color: &color.RGBA{R: 255, A: 255}}
	plot.Title.Text = "SAPAN GRAPH"
	plot.X.Label.Text = "request"
	plot.Y.Label.Text = "latency"

	//plot.Add(ScatterForData(data))
	//plot.Add(LineForAverage(data))
	plot.Add(Circle(5.0))

	wt, err := plot.WriterTo(512, 512, "png")
	if err != nil {
		return err
	}

	wt.WriteTo(w)
	return nil
}

// LineForAverage generates XYs for data
func LineForAverage(data []float64) *plotter.Line {
	xys := make(plotter.XYs, len(data))
	var sum float64 = 0.0

	for i, da := range data {
		sum += da
		xys[i].X = float64(i)
		xys[i].Y = sum / float64(i+1)
	}
	line, err := plotter.NewLine(xys)
	if err != nil {
		log.Printf("Could not plot line due to %v", err)
	}

	line.Color = color.RGBA{R: 255, A: 255}

	return line

}

// ScatterForData creates scatter with cross glyphs
func ScatterForData(data []float64) *plotter.Scatter {
	xys := make(plotter.XYs, len(data))

	for i, d := range data {
		xys[i].X = float64(i)
		xys[i].Y = d
	}

	scatter, err := plotter.NewScatter(xys)
	if err != nil {
		log.Printf("Could not create scatter due to %v", err)
	}
	scatter.GlyphStyle.Shape = draw.CrossGlyph{}
	return scatter
}

//Circle makes circle
func Circle(radius float64) *plotter.Polygon {
	xys := make(plotter.XYs, 10/0.01+1)
	var x, y float64
	var i int32 = 0
	for x = 0.0; x <= 10.0; x += 1.0 {
		y = math.Sqrt(math.Pow(radius, 2.0)-math.Pow(x-5, 2.0)) + 5
		log.Printf("x=%f y=%f rad=%f", x, y, radius)
		xys[i].X = x
		xys[i].Y = y
		i++
	}
	scatter, err := plotter.NewPolygon(xys)
	if err != nil {
		log.Printf("Could not create Circle scatter due to %v", err)
	}
	return scatter
}
