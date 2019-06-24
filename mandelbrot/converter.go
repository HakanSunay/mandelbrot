package mandelbrot

import (
	"image"
	"image/color"
	"sync"
)

type Converter struct {
	Picture   Picture
	Bounds    Bound
	Algorithm Algorithm
}

func (c* Converter) ExportImage() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, c.GetWidth(), c.GetHeight()))
	colorMatrix := c.GetPixelMatrix()
	for i := 0; i < c.GetWidth(); i++ {
		for j := 0; j < c.GetHeight(); j++ {
			img.Set(i, j, colorMatrix[i][j])
		}
	}
	return img
}

func (c *Converter) PixelToComplex(x, y int) complex128 {
	return c.Bounds.pixelToComplex(x, y, c.Picture)
}

func (c *Converter) ComputeIterations(num complex128) uint8 {
	return c.Algorithm.getIterations(num)
}

func (c *Converter) MaxIterations() uint8 {
	return c.Algorithm.MaxIterations
}

func (c* Converter) StartComputation(w *sync.WaitGroup, channel *chan int) {
	for n := 0; n < c.GetWorkerCount(); n++ {
		w.Add(1)
		go c.ComputeColumn(w, channel)
	}
}

func (c* Converter) ComputeColumn(w *sync.WaitGroup, channel *chan int) {
	pixelMatrix := c.GetPixelMatrix()
	for x := range *channel {
		for y := 0; y < c.GetHeight(); y++ {
			complexNumber := c.PixelToComplex(x, y)
			if iterations := c.ComputeIterations(complexNumber); iterations < c.MaxIterations() {
				pixelMatrix[x][y] = color.NRGBA{iterations * 255, iterations * 50, iterations * 20, 255}
			} else {
				pixelMatrix[x][y] = color.NRGBA{0, 0, 0, 255}
			}
		}
	}
	w.Done()
}

func (c *Converter) GetHeight() int {
	return c.Picture.Height
}

func (c *Converter) GetPixelMatrix() [][]color.NRGBA {
	return c.Picture.PixelMatrix
}

func (c *Converter) GetWorkerCount() int {
	return c.Algorithm.Workers
}

func (c *Converter) GetWidth() int {
	return c.Picture.Width
}