package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type student struct {
		firstName string
		lastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type class struct {
		students []student
	}
	// TODO: declare a map of modules being attended by multiple classes
	type module struct {
		moduleNumber int
		classes      []class
	}

	// TODO: output everything using fmt.Println()
	myStudent := student{
		firstName: "Nando",
		lastName:  "Schuermann",
	}

	myClass := class{
		students: []student{
			myStudent, myStudent, myStudent,
		},
	}

	myModule := module{
		moduleNumber: 346,
		classes: []class{
			myClass, myClass, myClass,
		},
	}
	fmt.Printf("Student: %v\nClass: %v\nModule: %v", myStudent, myClass, myModule)
}
