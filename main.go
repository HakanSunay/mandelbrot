package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"mandelbrot/mandelbrot"
	"os"
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

	var workers = *tasks
	var width, height = mandelbrot.GetDimensions(*dimension)
	var realMin, realMax, imagMin, imagMax = mandelbrot.GetRanges(*ranges)
	var fileName = *outputFile
	var complexity = *complexity
	var iterations = uint8(*iterations)

	bound := mandelbrot.Bound{realMin, realMax, imagMin, imagMax}
	picture := mandelbrot.Picture{width, height}
	algorithm := mandelbrot.Algorithm{complexity, iterations}
	engine := mandelbrot.Converter{picture, bound, algorithm}

	pixels := mandelbrot.CreatePixelMatrix(height, width)

	c := make(chan int, width)
	var w sync.WaitGroup

	start := time.Now()

	for n := 0; n < workers; n++ {
		w.Add(1)
		go mandelbrot.CalculateColumn(&w, &c, height, &engine, pixels)
	}

	for i := 0; i < width; i++ {
		c <- i
	}

	close(c)
	w.Wait()

	fmt.Println(time.Since(start))

	resultFile := image.NewNRGBA(image.Rect(0, 0, width, height))
	picture.PopulateImage(resultFile, pixels)

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if err = png.Encode(f, resultFile); err != nil {
		fmt.Println(err)
	}
}
