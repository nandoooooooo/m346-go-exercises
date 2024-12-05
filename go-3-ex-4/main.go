package main

import "fmt"

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = '\u2465'
	Seven = '\u2466'
	Eight = '\u2467'
	Nine  = '\u2468'
	Ten   = '\u2469'
	Jack  = 'J'
	Queen = 'Q'
	King  = 'K'
	Ace   = 'A'
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	colors := map[rune]string{
		Diamonds: "\u001B[31m",
		Spades:   "\u001B[30m",
		Clubs:    "\u001B[30m",
		Hearts:   "\u001B[31m",
	}
	resetColor := "\u001B[0m"

	// TODO: Loop over suits and ranks to output all combinations.
	for _, rank := range ranks {
		for _, suit := range suits {
			// White Background so you are able to see the black symbols
			fmt.Printf("%s\u001B[47m%c%s\u001B[47m%c\t\t", colors[suit], suit, resetColor, rank)
		}
		fmt.Println()
	}
}
