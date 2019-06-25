package mandelbrot

import "testing"

type differencePair struct {
	min, max float64
	result float64
}

func TestBound_realDif_TableDriven(t *testing.T){
	tableTests := [4]differencePair{{min: -2, max: 2, result: 4},
		{min: -3, max: 3, result: 6},
		{min: -10, max: 5, result: 15}}

	for _, tt := range tableTests{
		b := Bound{RealMin: tt.min, RealMax: tt.max}
		if b.realDif() != tt.result{
			t.Errorf("got %f, expected %f", b.realDif(), tt.result)
		}
	}
}

func TestBound_imagDif_TableDriven(t *testing.T){
	tableTests := [4]differencePair{{min: -2, max: 2, result: 4},
		{min: -3.2, max: 3.8, result: 7},
		{min: -10, max: 10, result: 20}}

	for _, value := range tableTests{
		b := Bound{ImagMin: value.min, ImagMax: value.max}
		if b.imagDif() != value.result{
			t.Errorf("got %f, expected %f", b.imagDif(), value.result)
		}
	}
}

func TestPicture_pixelToComplex_withMaxValuesExpectMaxBounds(t *testing.T) {
	p := Picture{Width: 640, Height: 480}
	b := Bound{-2, 2, -2, 2}
	if r:= p.pixelToComplex(p.Width,p.Height, b); r != complex(2,2){
		t.Errorf("got %f, expected %f", r, complex(2,2))
	}

}

func TestPicture_pixelToComplex_withMinValuesExpectMinBounds(t *testing.T) {
	p := Picture{Width: 640, Height: 480}
	b := Bound{-2, 2, -2, 2}
	if r:= p.pixelToComplex(0,0, b); r != complex(-2,-2){
		t.Errorf("got %f, expected %f", r, complex(-2,-2))
	}
}

func TestAlgorithm_getIterationsForCoordinateOutsideMandelbrotSet(t *testing.T) {
	a := Algorithm{Complexity: 8, MaxIterations: 50}
	if iters := a.getIterations(complex(2,2)); iters >= 50{
		t.Errorf("got %d iterations, expected to be less than 50", iters)
	}
}

func TestAlgorithm_getIterationsForCoordinateInMandelbrotSet(t *testing.T) {
	a := Algorithm{Complexity: 8, MaxIterations: 50}
	if iters := a.getIterations(complex(1,1)); iters < 50{
		t.Errorf("got %d iterations, expected to be 50", iters)
	}
}