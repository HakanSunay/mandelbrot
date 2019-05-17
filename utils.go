package main

import (
	"image/color"
	"mandelbrot/src"
	"strconv"
	"strings"
	"sync"
)

func calculateColumn(w *sync.WaitGroup, c *chan int, height int, converter *src.Converter, pixels [][]color.NRGBA) {
	for x := range *c {
		for y := 0; y < height; y++ {
			complexNumber := converter.PixelToComplex(x, y)
			if iterations := converter.Compute(complexNumber); iterations < converter.MaxIterations() {
				pixels[x][y] = color.NRGBA{iterations * 255, iterations * 50, iterations * 20, 255}
			} else {
				pixels[x][y] = color.NRGBA{0, 0, 0, 255}
			}
		}
	}
	w.Done()
}

func createPixelMatrix(h int, w int) [][]color.NRGBA {
	pixels := make([][]color.NRGBA, w)
	for r := range pixels {
		pixels[r] = make([]color.NRGBA, h)
	}
	return pixels
}

func getRanges(s string) (float64, float64, float64, float64) {
	ranges := strings.Split(s, ":")
	if len(ranges) != 4 {
		return 0, 0, 0, 0
	} else {
		rmin, _ := strconv.ParseFloat(ranges[0], 64)
		rmax, _ := strconv.ParseFloat(ranges[1], 64)
		imin, _ := strconv.ParseFloat(ranges[2], 64)
		imax, _ := strconv.ParseFloat(ranges[3], 64)
		return rmin, rmax, imin, imax
	}

}

func getDimensions(s string) (int, int) {
	dimensions := strings.Split(s, "x")
	if len(dimensions) != 2 {
		return 0, 0
	} else {
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])
		return width, height
	}
}
