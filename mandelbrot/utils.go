package mandelbrot

import (
	"image/color"
	"strconv"
	"strings"
)

const (
	DefaultWIDTH   = 640
	DefaultHEIGHT  = 480
	DefaultImagMin = -2
	DefaultImagMax = 2
	DefaultRealMin = -2
	DefaultRealMax = 2
)

func CreatePixelMatrix(h int, w int) [][]color.NRGBA {
	pixels := make([][]color.NRGBA, w)
	for r := range pixels {
		pixels[r] = make([]color.NRGBA, h)
	}
	return pixels
}

func GetRanges(s string) (float64, float64, float64, float64) {
	ranges := strings.Split(s, ":")
	rmin, err := strconv.ParseFloat(ranges[0], 64)
	if err != nil {
		rmin = DefaultRealMin
	}
	rmax, err := strconv.ParseFloat(ranges[1], 64)
	if err != nil {
		rmax = DefaultRealMax
	}
	imin, err := strconv.ParseFloat(ranges[2], 64)
	if err != nil {
		imin = DefaultImagMin
	}
	imax, err := strconv.ParseFloat(ranges[3], 64)
	if err != nil {
		imax = DefaultImagMax
	}
	return rmin, rmax, imin, imax

}

func GetDimensions(s string) (int, int) {
	dimensions := strings.Split(s, "x")
	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		width = DefaultHEIGHT
	}
	height, err := strconv.Atoi(dimensions[1])
	if err != nil {
		height = DefaultWIDTH
	}
	return width, height
}

func FillChannelWithColumns(c *chan int, width int) {
	for i := 0; i < width; i++ {
		*c <- i
	}
}