package main

import "fmt"

func main() {
	// TODO: create a map called "modules"

	modules := map[int]string{104: "Implement data models", 117: "Realise net infrastructure for a small company", 346: "Realise cloud solutions"}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117)
	// TODO: add one
	modules[320] = "OOP"
	// TODO: replace one
	modules[320] = "Teacher hopefully changes in close future"
	fmt.Println(modules)
}
