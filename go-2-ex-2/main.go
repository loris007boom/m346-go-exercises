package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[1] + fibs[0]
	fibs[3] = fibs[2] + fibs[1]
	fibs[4] = fibs[3] + fibs[2]

	next := fibs[3] + fibs[4]
	fibs = append(fibs, next)

	for i := 0; i < 3; i++ {
		n := fibs[len(fibs)-1] + fibs[len(fibs)-2]
		fibs = append(fibs, n)
	}

	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
