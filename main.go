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

const (
	THREADS    = 1
	DIMENSIONS = "640x480"
	RANGE      = "-2.0:2.0:-1.0:1.0"
	OUTPUTFILE = "zad18.png"
	COMPLEXITY = 8
	ITERATIONS = 50
)

var (
	tasks, iterations *int
	dimension,
	ranges,
	outputFile *string
	complexity *float64
)

func init() {
	tasks = flag.Int("t", THREADS, "Amount of threads")
	dimension = flag.String("s", DIMENSIONS, "Dimensions: width x height")
	ranges = flag.String("r", RANGE, "Real and Imaginary Number Range")
	outputFile = flag.String("o", OUTPUTFILE, "Name of the result file")
	complexity = flag.Float64("c", COMPLEXITY, "Fractal complexity")
	iterations = flag.Int("i", ITERATIONS, "Mandelbrot loop maximum iterations")
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
	bound := mandelbrot.NewBound(realMin, realMax, imagMin, imagMax)
	picture := mandelbrot.NewPicture(width, height, pixelMatrix)
	algorithm := mandelbrot.NewAlgorithm(complexity, iterations, workers)
	generator := mandelbrot.NewFractalGenerator(picture, bound, algorithm)

	c := make(chan int, width)
	var w sync.WaitGroup

	start := time.Now()
	generator.StartComputation(&w, &c)
	mandelbrot.FillChannelWithColumns(&c, width)

	close(c)
	w.Wait()

	parallelWorkTime := time.Since(start)
	defer fmt.Println(parallelWorkTime)

	coloredImage := generator.ExportImage()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if err := png.Encode(file, coloredImage); err != nil {
		fmt.Println(err)
	}
}
