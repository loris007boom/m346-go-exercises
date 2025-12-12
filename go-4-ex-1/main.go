package main

import (
	"errors"
	"fmt"
)

func computeGrade(gotPoints, maxPoints float64) (float64, error) {
	if maxPoints <= 0 {
		return 0, errors.New("maxPoints must be > 0")
	}
	if gotPoints < 0 {
		return 0, errors.New("gotPoints must be >= 0")
	}
	if gotPoints > maxPoints {
		return 0, errors.New("gotPoints cannot be greater than maxPoints")
	}

	grade := 1.0 + 5.0*(gotPoints/maxPoints)

	if grade < 1.0 {
		grade = 1.0
	}
	if grade > 6.0 {
		grade = 6.0
	}

	return grade, nil
}

func main() {
	g1, err := computeGrade(17.5, 28.0) // 4.125
	fmt.Println(g1, err)

	g2, err := computeGrade(28.0, 28.0) // 6
	fmt.Println(g2, err)

	g3, err := computeGrade(0.0, 28.0) // 1
	fmt.Println(g3, err)

	g4, err := computeGrade(30.0, 28.0) // error
	fmt.Println(g4, err)
}
