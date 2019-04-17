package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

const (
	MAX_ITERATIONS = 50
	REAL_MIN       = -2.
	REAL_MAX       = 2.
	IMAG_MIN       = -2.
	IMAG_MAX       = 2.
	WIDTH          = 2048
	HEIGHT         = 2048
	//WORKERS        = 4
	COMPLEXITY     = 8
)

var (
	tasks *int
	//dimension,
	//ranges,
	//outputFile *string
)

func init(){
	tasks = flag.Int("t", 1, "Amount of threads")
	//dimension = flag.String("s", "640x480", "Dimensions: width x height")
	//ranges = flag.String("r", "-2.0:2.0:-2.0:2.0", "Real and Imaginary Number Range" )
	//outputFile = flag.String("o", "zad18.png", "Name of the result file")
}

func main() {
	flag.Parse()
	var WORKERS = *tasks
	start := time.Now()

	bounds := image.Rect(0, 0, WIDTH, HEIGHT)
	resultFile := image.NewNRGBA(bounds)
	draw.Draw(resultFile, bounds, image.NewUniform(color.Black), image.ZP, draw.Src)
	
	c := make(chan int, WIDTH)

	var w sync.WaitGroup
	for n := 0; n < WORKERS; n++ {
		w.Add(1)
		go func() {
			for x := range c {
				for y := 0; y < HEIGHT; y++ {
					complexNumber := pixelToComplex(x, y)
					if iterations := mandelbrot(complexNumber); iterations < MAX_ITERATIONS {
						resultFile.Set(x, y, color.NRGBA{iterations * 255, iterations * 50, 20 * iterations, 255})
					}

				}
			}
			w.Done()
		}()
	}

	for i := 0; i < WIDTH; i++ {
		c <- i
	}
	close(c)
	w.Wait()

	f, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if err = png.Encode(f, resultFile); err != nil {
		fmt.Println(err)
	}
	fmt.Println(time.Since(start))
}

func mandelbrot(num complex128) uint8 {
	currentIterations := uint8(0)
	for z := num; cmplx.Abs(z) <= COMPLEXITY && currentIterations < MAX_ITERATIONS; currentIterations++ {
		z = cmplx.Cos(z) * num
	}
	return currentIterations
}

func pixelToComplex(x, y int) complex128 {
	return complex(REAL_MIN+float64(x)/WIDTH*(REAL_MAX - REAL_MIN),
		IMAG_MIN+float64(y)/HEIGHT*(IMAG_MAX - IMAG_MIN))
}