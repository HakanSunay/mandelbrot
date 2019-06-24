package mandelbrot

import (
	"image/color"
	"math/cmplx"
	"strconv"
	"strings"
)

type Bound struct {
	RealMin, RealMax, ImagMin, ImagMax float64
}

type Picture struct {
	Width, Height int
	PixelMatrix [][]color.NRGBA
}

type Algorithm struct {
	Complexity    float64
	MaxIterations uint8
	Workers int
}

func (b *Bound) RealDif() float64 {
	return b.RealMax - b.RealMin
}

func (b *Bound) ImagDif() float64 {
	return b.ImagMax - b.ImagMin
}

func (b *Bound) pixelToComplex(x, y int, picture Picture) complex128 {
	width := float64(picture.Width)
	height := float64(picture.Height)

	return complex(b.RealMin+(float64(x)/width)*b.RealDif(),
		b.ImagMin+(float64(y)/height)*b.ImagDif())
}

func (a *Algorithm) getIterations(num complex128) uint8 {
	currentIterations := uint8(0)
	for z := num; cmplx.Abs(z) <= a.Complexity && currentIterations < a.MaxIterations; currentIterations++ {
		z = cmplx.Cos(z) * num
	}
	return currentIterations
}

func CreatePixelMatrix(h int, w int) [][]color.NRGBA {
	pixels := make([][]color.NRGBA, w)
	for r := range pixels {
		pixels[r] = make([]color.NRGBA, h)
	}
	return pixels
}

func GetRanges(s string) (float64, float64, float64, float64) {
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

func GetDimensions(s string) (int, int) {
	dimensions := strings.Split(s, "x")
	if len(dimensions) != 2 {
		return 0, 0
	} else {
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])
		return width, height
	}
}

func FillChannelWithColumns(c *chan int, width int){
	for i := 0; i < width; i++ {
		*c <- i
	}
}