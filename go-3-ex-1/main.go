package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputWithZodiacSign(p Person) {
	var zodiacSign rune = '?'

	day := p.Day
	month := p.Month

	if (month == 3 && day >= 21) || (month == 4 && day <= 20) {
		zodiacSign = Aries
	} else if (month == 4 && day >= 21) || (month == 5 && day <= 21) {
		zodiacSign = Taurus
	} else if (month == 5 && day >= 22) || (month == 6 && day <= 21) {
		zodiacSign = Gemini
	} else if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		zodiacSign = Cancer
	} else if (month == 7 && day >= 23) || (month == 8 && day <= 23) {
		zodiacSign = Leo
	} else if (month == 8 && day >= 24) || (month == 9 && day <= 23) {
		zodiacSign = Virgo
	} else if (month == 9 && day >= 24) || (month == 10 && day <= 23) {
		zodiacSign = Libra
	} else if (month == 10 && day >= 24) || (month == 11 && day <= 22) {
		zodiacSign = Scorpius
	} else if (month == 11 && day >= 23) || (month == 12 && day <= 21) {
		zodiacSign = Sagittarius
	} else if (month == 12 && day >= 22) || (month == 1 && day <= 20) {
		zodiacSign = Capricornus
	} else if (month == 1 && day >= 21) || (month == 2 && day <= 19) {
		zodiacSign = Aquarius
	} else if (month == 2 && day >= 20) || (month == 3 && day <= 20) {
		zodiacSign = Pisces
	}

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

type FullName struct {
	FirstName string
	LastName  string
}
type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Person struct {
	FullName
	BirthDate
}

func main() {
	grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
	dennis := Person{FullName{"Dennis", "Ritchie"}, BirthDate{9, 9, 1941}}
	rick := Person{FullName{"Rick", "Astley"}, BirthDate{6, 2, 1966}}
	edsger := Person{FullName{"Edsger", "Dijkstra"}, BirthDate{11, 5, 1930}}
	alan := Person{FullName{"Alan", "Turing"}, BirthDate{23, 6, 1912}}

	outputWithZodiacSign(grace)
	outputWithZodiacSign(dennis)
	outputWithZodiacSign(rick)
	outputWithZodiacSign(edsger)
	outputWithZodiacSign(alan)

	me := Person{FullName{"Loris", "F"}, BirthDate{3, 5, 2005}}
	einstein := Person{FullName{"Albert", "Einstein"}, BirthDate{14, 3, 1879}}
	bill := Person{FullName{"Bill", "Gates"}, BirthDate{28, 10, 1955}}

	outputWithZodiacSign(me)
	outputWithZodiacSign(einstein)
	outputWithZodiacSign(bill)
}
