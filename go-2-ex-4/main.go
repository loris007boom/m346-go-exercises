package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Name     string
	Students []Student
}

type Module struct {
	Number  int
	Classes []Class
}

func main() {
	classA := Class{
		Name: "IM23a",
		Students: []Student{
			{FirstName: "Thomas", LastName: "Müller"},
			{FirstName: "Anna", LastName: "Muster"},
			{FirstName: "Tim", LastName: "Beispiel"},
		},
	}

	classB := Class{
		Name: "IM23b",
		Students: []Student{
			{FirstName: "Sara", LastName: "Meier"},
			{FirstName: "Luca", LastName: "Keller"},
			{FirstName: "Nina", LastName: "Winter"},
		},
	}

	classes := []Class{classA, classB}

	modules := map[int]Module{
		101: {
			Number:  101,
			Classes: []Class{classA},
		},
		117: {
			Number:  117,
			Classes: []Class{classA, classB},
		},
		346: {
			Number:  346,
			Classes: []Class{classB},
		},
	}

	fmt.Println("Klassen und Schüler:")
	for _, c := range classes {
		fmt.Println("Klasse:", c.Name)
		for _, s := range c.Students {
			fmt.Printf("  - %s %s\n", s.FirstName, s.LastName)
		}
	}

	fmt.Println("\nModule und teilnehmende Klassen:")
	for _, m := range modules {
		fmt.Printf("Modul %d wird besucht von:\n", m.Number)
		for _, c := range m.Classes {
			fmt.Printf("  - %s\n", c.Name)
		}
	}
}
