package main

import "fmt"

func main() {
	// declare string variable
	// and then assign string value into it.
	// var is short for variable
	var card string = newCard()
	var cards []string = []string{
		newCard(), newCard(),
	}
	strings := append(cards, card) // this does not modify the original slices.
	fmt.Println(cards)
	fmt.Println(strings)

	for i, card := range strings {
		println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
