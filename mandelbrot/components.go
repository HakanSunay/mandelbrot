package mandelbrot

import (
	"image/color"
	"math/cmplx"
)

type Bound struct {
	realMin, realMax, imagMin, imagMax float64
}

type Picture struct {
	width, height int
	pixelMatrix   [][]color.NRGBA
}

type Algorithm struct {
	complexity    float64
	maxIterations uint8
	workers       int
}

func NewBound(realMin, realMax, imagMin, imagMax float64) Bound{
	return Bound{realMin: realMin, realMax: realMax, imagMin: imagMin, imagMax: imagMax}
}

func (b *Bound) realDif() float64 {
	return b.realMax - b.realMin
}

func (b *Bound) imagDif() float64 {
	return b.imagMax - b.imagMin
}

func NewPicture(width, height int, pixelMatrix [][]color.NRGBA) Picture{
	return Picture{width:width, height:height, pixelMatrix:pixelMatrix}
}

func (p *Picture) pixelToComplex(x, y int, b Bound) complex128 {
	width := float64(p.width)
	height := float64(p.height)

	return complex(b.realMin+(float64(x)/width)*b.realDif(),
		b.imagMin+(float64(y)/height)*b.imagDif())
}

func NewAlgorithm(complexity float64, maxIterations uint8, workers int) Algorithm{
	return Algorithm{complexity:complexity, maxIterations:maxIterations, workers:workers}
}

func (a *Algorithm) getIterations(num complex128) uint8 {
	currentIterations := uint8(0)
	for z := num; cmplx.Abs(z) <= a.complexity && currentIterations < a.maxIterations; currentIterations++ {
		z = cmplx.Cos(z) * num
	}
	return currentIterations
}
