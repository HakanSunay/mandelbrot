package mandelbrot

import (
	"testing"
)

func TestGetRangesWithInvalidInputExpectDefaultValues(t *testing.T) {
	incorrectStrToBeTested := "asdasd:"
	if r1, r2, i1, i2 := GetRanges(incorrectStrToBeTested); r1 != DefaultRealMin ||
		r2 != DefaultRealMax || i1 != DefaultImagMin || i2 != DefaultImagMax {
		t.Errorf("expected default values for incorrect input %s", incorrectStrToBeTested)
	}
}

func TestGetRangesWithCorrectRangeExpectCorrectConversion(t *testing.T) {
	strToBeTested := "-1.5:2.5:-3.2:3.7"
	expectedRealMin := -1.5
	expectedRealMax := 2.5
	expectedImagMin := -3.2
	expectedImagMax := 3.7
	if r1, r2, i1, i2 := GetRanges(strToBeTested); r1 != expectedRealMin ||
		r2 != expectedRealMax || i1 != expectedImagMin || i2 != expectedImagMax {
		t.Errorf("expected %f:%f:%f:%f, got %f:%f:%f:%f",
			expectedRealMin, expectedRealMax, expectedImagMin, expectedImagMax,
			r1, r2, i1, i2)
	}
}

func TestCreatePixelMatrixForZeroHeightAndWidthExpectEmptyMatrix(t *testing.T) {
	if result := CreatePixelMatrix(0, 0); len(result) != 0 {
		t.Errorf("expected [] with length 0, but got length %d", len(result))
	}
}

func TestGetDimensionsWithIncorrectInputExpectDefaultValues(t *testing.T) {
	incorrectDimensionString := "imnotanumber"
	if w, h := GetDimensions(incorrectDimensionString); w != DefaultWIDTH || h != DefaultHEIGHT {
		t.Errorf("expected %dx%d, got %dx%d", DefaultWIDTH, DefaultHEIGHT, w, h)
	}
}

func TestGetDimensionsWithSemiCorrectInputExpectFirstInputAsResultAndDefault(t *testing.T) {
	semiCorrectDimensionString := "650x4890as"
	expectedWidth := 650
	if w, h := GetDimensions(semiCorrectDimensionString); w != expectedWidth || h != DefaultHEIGHT {
		t.Errorf("expected %dx%d, got %dx%d", expectedWidth, DefaultHEIGHT, w, h)
	}
}

func TestGetDimensionsWithCorrectInput(t *testing.T) {
	correctDimensionString := "650x2400"
	expectedWidth := 650
	expectedHeight := 2400
	if w, h := GetDimensions(correctDimensionString); w != expectedWidth || h != expectedHeight {
		t.Errorf("expected %dx%d, got %dx%d", expectedWidth, expectedHeight, w, h)
	}
}

func TestIndexDoesntExistExpectTrue(t *testing.T) {
	stringLength := 5
	indexNum := 5
	if !(indexDoesntExist(stringLength, indexNum)) {
		t.Errorf("expected true, got false")
	}
}

func TestIndexDoesntExistExpectFalse(t *testing.T) {
	stringLength := 5
	indexNum := 3
	if indexDoesntExist(stringLength, indexNum) {
		t.Errorf("expected false, got true")
	}
}
