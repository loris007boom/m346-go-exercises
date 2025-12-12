package main

import "fmt"

func convertCelsiusToFahrenheit(c float64) float64 {
	return (9.0/5.0)*c + 32.0
}

func convertFahrenheitToCelsius(f float64) float64 {
	return (5.0 / 9.0) * (f - 32.0)
}

type Celsius float64
type Fahrenheit float64

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit((9.0/5.0)*float64(c) + 32.0)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((5.0 / 9.0) * (float64(f) - 32.0))
}

func main() {
	f1 := convertCelsiusToFahrenheit(0.0)  // 32
	f2 := convertCelsiusToFahrenheit(23.0) // 73.4
	f3 := convertCelsiusToFahrenheit(100)  // 212
	fmt.Println(f1, f2, f3)

	fmt.Println(convertFahrenheitToCelsius(f1)) // 0
	fmt.Println(convertFahrenheitToCelsius(f2)) // 23
	fmt.Println(convertFahrenheitToCelsius(f3)) // 100

	var cozy Celsius = 23.0
	fmt.Println(cozy.ConvertToFahrenheit())               // 73.4
	fmt.Println(cozy.ConvertToFahrenheit().ConvertToCelsius()) // 23

	var cold Fahrenheit = 15.3
	fmt.Println(cold.ConvertToCelsius()) // -9.277777...

	// Zusatzfrage (Notation): beides geht, Methoden sind oft lesbarer bei Ketten.
	// fmt.Println(convertFahrenheitToCelsius(convertCelsiusToFahrenheit(23.0)))
	// var c Celsius = 23
	// fmt.Println(c.ConvertToFahrenheit().ConvertToCelsius())
}
