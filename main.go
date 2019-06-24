package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"rsa/mandelbrot"
	"sync"
	"time"
)

var (
	tasks, iterations *int
	dimension,
	ranges,
	outputFile *string
	complexity *float64
)

func init() {
	tasks = flag.Int("t", 1, "Amount of threads")
	dimension = flag.String("s", "640x480", "Dimensions: width x height")
	ranges = flag.String("r", "-2.0:2.0:-1.0:1.0", "Real and Imaginary Number Range")
	outputFile = flag.String("o", "zad18.png", "Name of the result file")
	complexity = flag.Float64("c", 8, "Fractal complexity")
	iterations = flag.Int("i", 50, "Mandelbrot loop maximum iterations")
}

func main() {
	flag.Parse()

	var (
		workers                            = *tasks
		width, height                      = mandelbrot.GetDimensions(*dimension)
		realMin, realMax, imagMin, imagMax = mandelbrot.GetRanges(*ranges)
		fileName                           = *outputFile
		complexity                         = *complexity
		iterations                         = uint8(*iterations)
	)

	pixelMatrix := mandelbrot.CreatePixelMatrix(height, width)
	bound := mandelbrot.Bound{RealMin: realMin, RealMax: realMax, ImagMin: imagMin, ImagMax: imagMax}
	picture := mandelbrot.Picture{Width: width, Height: height, PixelMatrix: pixelMatrix}
	algorithm := mandelbrot.Algorithm{Complexity: complexity, MaxIterations: iterations, Workers: workers}
	engine := mandelbrot.Converter{Picture: picture, Bounds: bound, Algorithm: algorithm}

	c := make(chan int, width)
	var w sync.WaitGroup

	start := time.Now()
	engine.StartComputation(&w, &c)
	mandelbrot.FillChannelWithColumns(&c, width)

	close(c)
	w.Wait()

	parallelWorkTime := time.Since(start)
	defer fmt.Println(parallelWorkTime)

	coloredImage := engine.ExportImage()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if err := png.Encode(file, coloredImage); err != nil {
		fmt.Println(err)
	}
}
