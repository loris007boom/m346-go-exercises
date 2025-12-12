package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

type ShortSides struct {
	a float64
	b float64
}

func (s ShortSides) Hypotenuse() float64 {
	return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}

func main() {
	fmt.Println(computeHypotenuse(3, 4))    // 5
	fmt.Println(computeHypotenuse(5, 12))   // 13
	fmt.Println(computeHypotenuse(8, 15))   // 17

	s1 := ShortSides{a: 3, b: 4}
	fmt.Println(s1.Hypotenuse()) // 5

	s2 := ShortSides{a: 5, b: 12}
	fmt.Println(s2.Hypotenuse()) // 13

	s3 := ShortSides{a: 8, b: 15}
	fmt.Println(s3.Hypotenuse()) // 17
}
