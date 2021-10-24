package onelab_homework1_lib

import (
	"errors"
	"fmt"
	"math"
)

var ErrNegativeDiscriminant = errors.New("the quadratic equation has no real roots")

func ChangeLetterCase(str string) string {
	arrRune := []rune(str)

	for i := range arrRune {
		if arrRune[i] >= 'A' && arrRune[i] <= 'Z' {
			arrRune[i] += 32
		} else if arrRune[i] >= 'a' && arrRune[i] <= 'z' {
			arrRune[i] -= 32
		}
	}

	return string(arrRune)
}

func RootsOfQuadraticEquation(a, b, c float64) (x1 float64, x2 float64, err error) {
	dc := b*b - 4*a*c

	switch {
	case dc > 0:
		x1 = (-b + math.Sqrt(dc)) / (2 * a)
		x2 = (-b - math.Sqrt(dc)) / (2 * a)
		return
	case dc == 0:
		x1 = -b / (2 * a)
		x2 = -b / (2 * a)
		return
	default:
		return x1, x2, fmt.Errorf("RootsOfQuadraticEquation: %w", ErrNegativeDiscriminant)
	}
}
