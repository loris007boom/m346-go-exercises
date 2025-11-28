package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	me := Profile{
		FullName: FullName{
			FirstName: "Loris",
			LastName:  "F",
		},
		BirthDate: BirthDate{
			Day:   3,
			Month: 5,
			Year:  2005,
		},
		NumberOfSiblings: 1,
		ZodiacSign: '♉',
	}

	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
