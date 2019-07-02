package mandelbrot

import (
	"image/color"
	"strconv"
	"strings"
)

const (
	DefaultWIDTH   = 640
	DefaultHEIGHT  = 480
	DefaultImagMin = -2.
	DefaultImagMax = 2.
	DefaultRealMin = -2.
	DefaultRealMax = 2.
)

// CreatePixelMatrix initializes a 2d array of type color.NRGBA
// for its height and width uses the user-given parameters
func CreatePixelMatrix(h int, w int) [][]color.NRGBA {
	pixels := make([][]color.NRGBA, w)
	for r := range pixels {
		pixels[r] = make([]color.NRGBA, h)
	}
	return pixels
}

// FillChannelWithRows loops through the numbers
// from 0 to width - 1 and sends them to a given channel
func FillChannelWithRows(c *chan int, width int) {
	for i := 0; i < width; i++ {
		*c <- i
	}
}

// GetRanges extracts the bounds from the input of the user
func GetRanges(s string) (float64, float64, float64, float64) {
	ranges := strings.Split(s, ":")
	rmin := parseFloat(ranges, 0, DefaultRealMin)
	rmax := parseFloat(ranges, 1, DefaultRealMax)
	imin := parseFloat(ranges, 2, DefaultImagMin)
	imax := parseFloat(ranges, 3, DefaultImagMax)
	return rmin, rmax, imin, imax
}

// GetDimensions extracts the width and the height from the input of the user
func GetDimensions(s string) (int, int) {
	dimensions := strings.Split(s, "x")
	width := int(parseFloat(dimensions, 0, DefaultWIDTH))
	height := int(parseFloat(dimensions, 1, DefaultHEIGHT))
	return width, height
}

// parseFloat is used to eliminate DRY when parsing numbers from string
// if the soughtIndex is out of bounds a default value is provided
func parseFloat(values []string, soughtIndexOfValue int, defaultValue float64) float64 {
	if indexDoesntExist(len(values), soughtIndexOfValue) {
		return defaultValue
	}
	result, err := strconv.ParseFloat(values[soughtIndexOfValue], 64)
	if err != nil {
		result = defaultValue
	}
	return result
}

// indexDoesntExist checks if the sought index is out of bounds
func indexDoesntExist(length int, index int) bool {
	if length < index+1 {
		return true
	}
	return false
}
