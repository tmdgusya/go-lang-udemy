package main

// Create a New type of deck
// which is a slices of string.
type deck []string

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}
