// Package mandelbrot provides mathematical abstractions
// and a user-defined generator to create a fractal
package mandelbrot

import (
	"image"
	"image/color"
	"sync"
)

const (
	RedColorAmplifier   = 255
	GreenColorAmplifier = 50
	BlueColorAmplifier  = 20
	AlphaAmplifier      = 255
)

type FractalGenerator struct {
	picture   Picture
	bounds    Bound
	algorithm Algorithm
}

func NewFractalGenerator(picture Picture, bound Bound, algorithm Algorithm) *FractalGenerator {
	return &FractalGenerator{picture: picture, bounds: bound, algorithm: algorithm}
}

func (fg *FractalGenerator) StartComputation(w *sync.WaitGroup, channel *chan int) {
	for n := 0; n < fg.getWorkerCount(); n++ {
		w.Add(1)
		go fg.computeAvailableRow(w, channel)
	}
}

func (fg *FractalGenerator) computeAvailableRow(w *sync.WaitGroup, channel *chan int) {
	pixelMatrix := fg.getPixelMatrix()
	for x := range *channel {
		for y := 0; y < fg.getHeight(); y++ {
			complexNumber := fg.pixelToComplex(x, y)
			iterations := fg.computeIterations(complexNumber)
			if fg.belongsToMandelbrotSet(iterations) {
				pixelMatrix[x][y] = color.NRGBA{A: AlphaAmplifier}
			} else {
				pixelMatrix[x][y] = color.NRGBA{R: iterations * RedColorAmplifier,
					G: iterations * GreenColorAmplifier,
					B: iterations * BlueColorAmplifier,
					A: AlphaAmplifier}
			}
		}
	}
	w.Done()
}

func (fg *FractalGenerator) ExportImage() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, fg.getWidth(), fg.getHeight()))
	colorMatrix := fg.getPixelMatrix()
	for i := 0; i < fg.getWidth(); i++ {
		for j := 0; j < fg.getHeight(); j++ {
			img.Set(i, j, colorMatrix[i][j])
		}
	}
	return img
}

func (fg *FractalGenerator) belongsToMandelbrotSet(iterations uint8) bool {
	if iterations < fg.maxIterations() {
		return false
	} else {
		return true
	}
}

func (fg *FractalGenerator) pixelToComplex(x, y int) complex128 {
	return fg.picture.pixelToComplex(x, y, fg.bounds)
}

func (fg *FractalGenerator) computeIterations(num complex128) uint8 {
	return fg.algorithm.getIterations(num)
}

func (fg *FractalGenerator) maxIterations() uint8 {
	return fg.algorithm.maxIterations
}

func (fg *FractalGenerator) getPixelMatrix() [][]color.NRGBA {
	return fg.picture.pixelMatrix
}

func (fg *FractalGenerator) getWorkerCount() int {
	return fg.algorithm.workers
}

func (fg *FractalGenerator) getWidth() int {
	return fg.picture.width
}

func (fg *FractalGenerator) getHeight() int {
	return fg.picture.height
}
