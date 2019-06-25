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

func FillChannelWithColumns(c *chan int, width int) {
	for i := 0; i < width; i++ {
		*c <- i
	}
}

func GetRanges(s string) (float64, float64, float64, float64) {
	ranges := strings.Split(s, ":")
	rmin := parseFloat(ranges, 0, DefaultRealMin)
	rmax := parseFloat(ranges, 1, DefaultRealMax)
	imin := parseFloat(ranges, 2, DefaultImagMin)
	imax := parseFloat(ranges, 3, DefaultImagMax)
	return rmin, rmax, imin, imax

}

func GetDimensions(s string) (int, int) {
	dimensions := strings.Split(s, "x")
	width := int(parseFloat(dimensions, 0, DefaultWIDTH))
	height := int(parseFloat(dimensions, 1, DefaultHEIGHT))
	return width, height
}

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

func indexDoesntExist(length int, index int) bool {
	if length < index+1 {
		return true
	}
	return false
}
