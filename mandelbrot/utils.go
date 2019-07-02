package mandelbrot

import (
	"image/color"
	"strconv"
	"strings"
)

const (
	DefaultWIDTH   = 640
	DefaultHEIGHT  = 480
	IndexOfWidth = 0
	IndexOfHeight = 1
	DimentionsDelimiter = "x"
	DefaultImagMin = -2.
	DefaultImagMax = 2.
	DefaultRealMin = -2.
	DefaultRealMax = 2.
	IndexOfRealMin = 0
	IndexOfRealMax = 1
	IndexOfImagMin = 2
	IndexOfImagMax = 3
	RangesDelimiter = ":"
	FloatBitSize = 64
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
	ranges := strings.Split(s, RangesDelimiter)
	rmin := parseFloat(ranges, IndexOfRealMin, DefaultRealMin)
	rmax := parseFloat(ranges, IndexOfRealMax, DefaultRealMax)
	imin := parseFloat(ranges, IndexOfImagMin, DefaultImagMin)
	imax := parseFloat(ranges, IndexOfImagMax, DefaultImagMax)
	return rmin, rmax, imin, imax
}

// GetDimensions extracts the width and the height from the input of the user
func GetDimensions(s string) (int, int) {
	dimensions := strings.Split(s, DimentionsDelimiter)
	width := int(parseFloat(dimensions, IndexOfWidth, DefaultWIDTH))
	height := int(parseFloat(dimensions, IndexOfHeight, DefaultHEIGHT))
	return width, height
}

// parseFloat is used to eliminate DRY when parsing numbers from string
// if the soughtIndex is out of bounds a default value is provided
func parseFloat(values []string, soughtIndexOfValue int, defaultValue float64) float64 {
	if indexDoesntExist(len(values), soughtIndexOfValue) {
		return defaultValue
	}
	result, err := strconv.ParseFloat(values[soughtIndexOfValue], FloatBitSize)
	if err != nil {
		result = defaultValue
	}
	return result
}

// indexDoesntExist checks if the sought index is out of bounds
func indexDoesntExist(length int, index int) bool {
	if length < index + 1 {
		return true
	}
	return false
}
