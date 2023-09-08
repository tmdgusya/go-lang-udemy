package main

func main() {
	// declare string variable
	// and then assign string value into it.
	// var is short for variable
	var cards deck = newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
