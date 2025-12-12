package main

import (
	"fmt"
	"math"
)

func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a, b, c float64) []float64 {
	if a == 0 {
		return []float64{}
	}

	D := computeDiscriminant(a, b, c)

	if D < 0 {
		return []float64{}
	}

	if D == 0 {
		x := (-b) / (2 * a)
		return []float64{x}
	}

	sqrtD := math.Sqrt(D)
	x1 := (-b + sqrtD) / (2 * a)
	x2 := (-b - sqrtD) / (2 * a)
	return []float64{x1, x2}
}

func main() {
	fmt.Println(computeDiscriminant(3, 4, 1)) // 4
	fmt.Println(computeDiscriminant(2, 4, 2)) // 0
	fmt.Println(computeDiscriminant(3, 4, 2)) // -8

	fmt.Println(computeQuadraticFormula(3, 4, 1)) // [-0.3333333333333333 -1]
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // [-1]
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // []
}
