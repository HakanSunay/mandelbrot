package main

import (
	"image"
	"image/color"
	"math/cmplx"
	"strconv"
	"strings"
	"sync"
)

type Converter struct {
	Width, Height                                  int
	RealMin, RealMax, ImagMin, ImagMax, Complexity float64
	MaxIterations                                  uint8
}

func NewConverter(i int, i2 int, f float64, f2 float64, f3 float64, f4 float64, f5 float64, i3 uint8) *Converter {
	return &Converter{i, i2, f, f2, f3,
		f4, f5, i3}
}

func (c *Converter) populateImage(img *image.NRGBA, colorMatrix [][]color.NRGBA) {
	for i := 0; i < c.Width; i++ {
		for j := 0; j < c.Height; j++ {
			img.Set(i, j, colorMatrix[i][j])
		}
	}
}

func (c *Converter) mandelbrot(num complex128) uint8 {
	currentIterations := uint8(0)
	for z := num; cmplx.Abs(z) <= c.Complexity && currentIterations < c.MaxIterations; currentIterations++ {
		z = cmplx.Cos(z) * num
	}
	return currentIterations
}

func (c *Converter) pixelToComplex(x, y int) complex128 {
	return complex(c.RealMin+(float64(x)/float64(c.Width))*(c.RealMax-c.RealMin),
		c.ImagMin+(float64(y)/float64(c.Height))*(c.ImagMax-c.ImagMin))
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

func calculateColumn(w *sync.WaitGroup, c *chan int, height int, converter *Converter, pixels [][]color.NRGBA) {
	for x := range *c {
		for y := 0; y < height; y++ {
			complexNumber := converter.pixelToComplex(x, y)
			if iterations := converter.mandelbrot(complexNumber); iterations < converter.MaxIterations {
				pixels[x][y] = color.NRGBA{iterations * 255, iterations * 50,
					iterations * 20, 255}
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
