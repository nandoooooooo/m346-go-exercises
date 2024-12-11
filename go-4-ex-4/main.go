package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
// TODO: implement the function convertFahrenheitToCelsius

type Celsius float64
type Fahrenheit float64

func main() {

	/*

		Zusatzfrage:

		Variante 2, weil man erstens mit einer Variable "c" arbeitet, man per methodchaining diese verändert und
		schlussendlich printed. Variante 1 ist nur eine lange Zeile.

	*/

	// TODO: call the function convertCelsiusToFahrenheit
	c1, c2, c3 := -2.0, 12.5, 51.2
	fmt.Printf("%v°C => %v°F\n", c1, convertCelsiusToFahrenheit(c1))   // 28.4 °F
	fmt.Printf("%v°C => %v°F\n", c2, convertCelsiusToFahrenheit(c2))   // 54.5 °F
	fmt.Printf("%v°C => %v°F\n\n", c3, convertCelsiusToFahrenheit(c3)) // 124.16 °F
	// TODO: call the function convertFahrenheitToCelsius
	f1, f2, f3 := 1.0, 100.0, 77.23
	fmt.Printf("%v°C => %v°F\n", f1, convertFahrenheitToCelsius(f1))   // -17.222 °C
	fmt.Printf("%v°C => %v°F\n", f2, convertFahrenheitToCelsius(f2))   // 37.777 °C
	fmt.Printf("%v°C => %v°F\n\n", f3, convertFahrenheitToCelsius(f3)) // 25.12777 °C

	var cozy Celsius = 23.0
	var cold Fahrenheit = 15.3

	fmt.Printf("Cozy: %v°C => %v°F\n", cozy, cozy.convertCelsiusToFahrenheitWithStruct()) // 73.4 °F
	fmt.Printf("Cold: %v°F => %v°C\n", cold, cold.convertFahrenheitToCelsiusWithStruct()) // -9.277 °C
}

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func (c Celsius) convertCelsiusToFahrenheitWithStruct() float64 {
	return float64(c)*9/5 + 32
}

func (f Fahrenheit) convertFahrenheitToCelsiusWithStruct() float64 {
	return (float64(f) - 32) * 5 / 9
}
