package onelab_homework1_lib

import (
	"errors"
	"testing"
)

func TestChangeLetterCase(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"alO", "ALo"},
		{"laLAlaA", "LAlaLAa"},
		{"aAaaaAAaaA", "AaAAAaaAAa"},
	}

	for _, tc := range testCases {
		if str := ChangeLetterCase(tc.input); str != tc.output {
			t.Errorf("ChangeLetterCase(%s) => got: %s, want: %s", tc.input, str, tc.output)
		}
	}
}

func TestRootsOfQuadraticEquation(t *testing.T) {
	testCases := []struct {
		a   float64
		b   float64
		c   float64
		x1  float64
		x2  float64
		err error
	}{
		{1, 4, 5, 0, 0, ErrNegativeDiscriminant},
		{2, 3, -5, 1, -2.5, nil},
		{1, 2, 1, -1, -1, nil},
	}

	for _, tc := range testCases {
		x1, x2, err := RootsOfQuadraticEquation(tc.a, tc.b, tc.c)
		if x1 != tc.x1 || x2 != tc.x2 || !errors.Is(err, tc.err) {
			t.Errorf("RootsOfQuadraticEquation(%f, %f, %f) => got (%f, %f, %s), want (%f, %f, %s)", tc.a, tc.b, tc.c, x1, x2, err, tc.x1, tc.x2, tc.err)
		}
	}
}
