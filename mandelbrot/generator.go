package mandelbrot

import (
	"image"
	"image/color"
	"sync"
)

type Generator struct {
	Picture   Picture
	Bounds    Bound
	Algorithm Algorithm
}

func (g * Generator) ExportImage() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, g.GetWidth(), g.GetHeight()))
	colorMatrix := g.GetPixelMatrix()
	for i := 0; i < g.GetWidth(); i++ {
		for j := 0; j < g.GetHeight(); j++ {
			img.Set(i, j, colorMatrix[i][j])
		}
	}
	return img
}

func (g *Generator) PixelToComplex(x, y int) complex128 {
	return g.Bounds.pixelToComplex(x, y, g.Picture)
}

func (g *Generator) ComputeIterations(num complex128) uint8 {
	return g.Algorithm.getIterations(num)
}

func (g *Generator) MaxIterations() uint8 {
	return g.Algorithm.MaxIterations
}

func (g * Generator) StartComputation(w *sync.WaitGroup, channel *chan int) {
	for n := 0; n < g.GetWorkerCount(); n++ {
		w.Add(1)
		go g.ComputeColumn(w, channel)
	}
}

func (g * Generator) ComputeColumn(w *sync.WaitGroup, channel *chan int) {
	pixelMatrix := g.GetPixelMatrix()
	for x := range *channel {
		for y := 0; y < g.GetHeight(); y++ {
			complexNumber := g.PixelToComplex(x, y)
			if iterations := g.ComputeIterations(complexNumber); iterations < g.MaxIterations() {
				pixelMatrix[x][y] = color.NRGBA{R: iterations * 255, G: iterations * 50, B: iterations * 20, A: 255}
			} else {
				pixelMatrix[x][y] = color.NRGBA{A: 255}
			}
		}
	}
	w.Done()
}

func (g *Generator) GetHeight() int {
	return g.Picture.Height
}

func (g *Generator) GetPixelMatrix() [][]color.NRGBA {
	return g.Picture.PixelMatrix
}

func (g *Generator) GetWorkerCount() int {
	return g.Algorithm.Workers
}

func (g *Generator) GetWidth() int {
	return g.Picture.Width
}