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

// Fractal Generator is a manager struct that helps us
// perform the main operations when creating the fractal
type FractalGenerator struct {
	picture   Picture
	bounds    Bound
	algorithm Algorithm
}

// NewFractalGenerator initializes a new FractalGenerator struct
// using the passed in components
func NewFractalGenerator(picture Picture, bound Bound, algorithm Algorithm) *FractalGenerator {
	return &FractalGenerator{picture: picture, bounds: bound, algorithm: algorithm}
}

// StartComputation generates the fractal image
// it uses a WaitGroup and a Channel to time and
// paralyze the mathematical computation of the mandelbrot set
func (fg *FractalGenerator) StartComputation(w *sync.WaitGroup, channel *chan int) {
	for n := 0; n < fg.getWorkerCount(); n++ {
		w.Add(1)
		go fg.computeAvailableRow(w, channel)
	}
}

// computeAvailableRow ranges the passed in channel and
// when given a certain number, it generates the full
// mandelbrot picture in the row corresponding to the given number
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

// ExportImage is used to export the generated image
// in the previous process started by StartComputation
// if it is called before StartComputation
// it will generate an empty image
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

// belongsToMandelbrotSet checks if the iterations of a given
// coordinate belongs to the Mandelbrot set
func (fg *FractalGenerator) belongsToMandelbrotSet(iterations uint8) bool {
	if iterations < fg.maxIterations() {
		return false
	} else {
		return true
	}
}

// pixelToComplex converts a pixel represented by x and y coordinates
// into a number in the Complex plane - having real and imaginary coordinates
func (fg *FractalGenerator) pixelToComplex(x, y int) complex128 {
	return fg.picture.pixelToComplex(x, y, fg.bounds)
}

// computeIterations calls the algorithm to get the iterations
// of a particular complex number
func (fg *FractalGenerator) computeIterations(num complex128) uint8 {
	return fg.algorithm.getIterations(num)
}

// maxIterations returns the limit beyond which a point
// is considered inside the Mandelbrot set
func (fg *FractalGenerator) maxIterations() uint8 {
	return fg.algorithm.maxIterations
}

// getPixelMatrix returns the image represented as a pixel matrix
func (fg *FractalGenerator) getPixelMatrix() [][]color.NRGBA {
	return fg.picture.pixelMatrix
}

// getWorkerCount provides us with the number of threads specified
// by the user of the app
func (fg *FractalGenerator) getWorkerCount() int {
	return fg.algorithm.workers
}

// getWidth provides us with the width of the desired image
func (fg *FractalGenerator) getWidth() int {
	return fg.picture.width
}

// getHeight provides us with the height of the desired image
func (fg *FractalGenerator) getHeight() int {
	return fg.picture.height
}
