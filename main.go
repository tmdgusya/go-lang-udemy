package main

import "fmt"

func main() {
	// declare string variable
	// and then assign string value into it.
	// var is short for variable
	var card string = newCard()
	var cards deck = deck{
		newCard(), newCard(),
	}
	strings := append(cards, card) // this does not modify the original slices.
	fmt.Println(cards)
	fmt.Println(strings)

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
