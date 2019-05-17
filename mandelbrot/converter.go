package mandelbrot

import (
	"image"
	"image/color"
	"math/cmplx"
)

type Bound struct {
	RealMin, RealMax, ImagMin, ImagMax float64
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

type Picture struct {
	Width, Height int
}

func (p *Picture) PopulateImage(img *image.NRGBA, colorMatrix [][]color.NRGBA) {
	for i := 0; i < p.Width; i++ {
		for j := 0; j < p.Height; j++ {
			img.Set(i, j, colorMatrix[i][j])
		}
	}
}

type Algorithm struct {
	Complexity    float64
	MaxIterations uint8
}

func (a *Algorithm) getIterations(num complex128) uint8 {
	currentIterations := uint8(0)
	for z := num; cmplx.Abs(z) <= a.Complexity && currentIterations < a.MaxIterations; currentIterations++ {
		z = cmplx.Cos(z) * num
	}
	return currentIterations
}

type Converter struct {
	Picture   Picture
	Bounds    Bound
	Algorithm Algorithm
}

func (c *Converter) PixelToComplex(x, y int) complex128 {
	return c.Bounds.pixelToComplex(x, y, c.Picture)
}

func (c *Converter) Compute(num complex128) uint8 {
	return c.Algorithm.getIterations(num)
}

func (c *Converter) MaxIterations() uint8 {
	return c.Algorithm.MaxIterations
}
