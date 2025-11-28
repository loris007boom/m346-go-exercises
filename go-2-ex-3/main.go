package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Präsentationen erstellen und durchführen",
		117: "Datenbanken implementieren",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)

	modules[101] = "Informatikgrundlagen erarbeiten"

	modules[346] = "Moderne Cloud-Lösungen planen und umsetzen"

	fmt.Println(modules)
}
