package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	firstName := "Nando"
	lastName := "Schuermann"
	dayOfBirth := 14
	monthOfBirth := 2
	yearOfBirth := 2008
	numberOfSiblings := 1
	heightInMeters := 1.83
	zodiacSign := '\u2652'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
