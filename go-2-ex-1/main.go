package main

import "fmt"

type FullName struct {
	// TODO: add fields
	firstName string
	lastName  string
}

// TODO: declare a structure for birth date
type birthDate struct {
	DayOfBirth   int
	MonthOfBirth int
	YearOfBirth  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	fullName         FullName
	birthDate        birthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		fullName: FullName{
			"Nando", "Schuermann",
		},
		birthDate: birthDate{
			14, 2, 2008,
		},
		NumberOfSiblings: 1,        // TODO: adjust
		ZodiacSign:       '\u2652', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
