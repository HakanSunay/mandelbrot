package mandelbrot

import (
	"testing"
)

func TestFractalGenerator_ExportImageExpectEmptyImageForNotComputedFractal(t *testing.T) {
	fractalGenerator := FractalGenerator{}
	if img := fractalGenerator.ExportImage(); len(img.Pix) != 0 {
		t.Errorf("expected an empty image, got %d", img)
	}
}

func TestFractalGenerator_belongsToMandelbrotSet(t *testing.T) {
	fractalGenerator := FractalGenerator{}
	fractalGenerator.algorithm.maxIterations = 50
	if fractalGenerator.belongsToMandelbrotSet(fractalGenerator.maxIterations() - 1) {
		t.Errorf("expected false, got true")
	}
}
