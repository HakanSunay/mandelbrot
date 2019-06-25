package mandelbrot

import (
	"image/color"
	"math/cmplx"
)

// TODO: Add tests

type Bound struct {
	RealMin, RealMax, ImagMin, ImagMax float64
}

type Picture struct {
	Width, Height int
	PixelMatrix   [][]color.NRGBA
}

type Algorithm struct {
	Complexity    float64
	MaxIterations uint8
	Workers       int
}

func (b *Bound) realDif() float64 {
	return b.RealMax - b.RealMin
}

func (b *Bound) imagDif() float64 {
	return b.ImagMax - b.ImagMin
}

func (p *Picture) pixelToComplex(x, y int, b Bound) complex128 {
	width := float64(p.Width)
	height := float64(p.Height)

	return complex(b.RealMin+(float64(x)/width)*b.realDif(),
		b.ImagMin+(float64(y)/height)*b.imagDif())
}

func (a *Algorithm) getIterations(num complex128) uint8 {
	currentIterations := uint8(0)
	for z := num; cmplx.Abs(z) <= a.Complexity && currentIterations < a.MaxIterations; currentIterations++ {
		z = cmplx.Cos(z) * num
	}
	return currentIterations
}
