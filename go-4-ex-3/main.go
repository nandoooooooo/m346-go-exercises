package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula

func main() {
	// TODO: call the function computeQuadraticFormula
	a1, b1, c1 := 3.0, 4.0, 1.0
	result1 := computeQuadraticFormula(a1, b1, c1)
	fmt.Printf("Solutions for a = %v, b = %v , c = %v:\n\n%v\n\n", a1, b1, c1, result1) // [-0.3333, -1]

	a2, b2, c2 := 2.0, 4.0, 2.0
	result2 := computeQuadraticFormula(a2, b2, c2)
	fmt.Printf("Solutions for a = %v, b = %v , c = %v:\n\n%v\n\n", a2, b2, c2, result2) // [-1]

	a3, b3, c3 := 3.0, 4.0, 2.0
	result3 := computeQuadraticFormula(a3, b3, c3)
	fmt.Printf("Solutions for a = %v, b = %v , c = %v:\n\n%v\n\n", a3, b3, c3, result3) // []
}

func computeQuadraticFormula(a, b, c float64) []float64 {
	D := computeDiscriminant(a, b, c)

	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		return []float64{x1, x2}
	} else if D == 0 {
		x := -b / (2 * a)
		return []float64{x}
	} else {
		return []float64{}
	}
}

func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}
