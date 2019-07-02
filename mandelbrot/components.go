package mandelbrot

import (
	"image/color"
	"math/cmplx"
)

// Bound is a sturct that
// represents a mathematical abstraction
// of the limits in which the fractal can reside
// the elements of this struct are specified by the user
type Bound struct {
	realMin, realMax, imagMin, imagMax float64
}

// Picture is a struct that holds
// information about the description of the picture
// and its colour tones represented as a color matrix
type Picture struct {
	width, height int
	pixelMatrix   [][]color.NRGBA
}

// Algorithm is a struct that
// represents a mathematical abstraction of the mandelbrot set
// its elements are user defined which provides more flexibity
// when generating more complex fractals
type Algorithm struct {
	complexity    float64
	maxIterations uint8
	workers       int
}

// NewBound initializes a new Bound struct
// using the passed in limitations by the user
func NewBound(realMin, realMax, imagMin, imagMax float64) Bound{
	return Bound{realMin: realMin, realMax: realMax, imagMin: imagMin, imagMax: imagMax}
}

// realDif calculates the difference between
// the upper and lower real number bound
func (b *Bound) realDif() float64 {
	return b.realMax - b.realMin
}

// imagDif calculates the difference between
// the upper and lower imaginary number bound
func (b *Bound) imagDif() float64 {
	return b.imagMax - b.imagMin
}

// NewPicture initializes a new Picture struct
// using the passed in dimensions and the automatically generated color matrix
func NewPicture(width, height int, pixelMatrix [][]color.NRGBA) Picture{
	return Picture{width:width, height:height, pixelMatrix:pixelMatrix}
}

// pixelToComplex converts a given pixel with its x and y coordinates
// into a complex number by scaling the coordinates using the bounds
func (p *Picture) pixelToComplex(x, y int, b Bound) complex128 {
	width := float64(p.width)
	height := float64(p.height)

	return complex(b.realMin+(float64(x)/width)*b.realDif(),
		b.imagMin+(float64(y)/height)*b.imagDif())
}

// NewAlgorithm initializes a new Algorithm struct that is user defined
func NewAlgorithm(complexity float64, maxIterations uint8, workers int) Algorithm{
	return Algorithm{complexity:complexity, maxIterations:maxIterations, workers:workers}
}

// getIterations computes the iterations of a given complex number
// by using the predefined fractal algorithm
func (a *Algorithm) getIterations(num complex128) uint8 {
	currentIterations := uint8(0)
	for z := num; cmplx.Abs(z) <= a.complexity && currentIterations < a.maxIterations; currentIterations++ {
		// the mandelbrot fractal generating algorithm
		z = cmplx.Cos(z) * num
	}
	return currentIterations
}
