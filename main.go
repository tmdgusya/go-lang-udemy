package main

import "fmt"

func main() {
	// declare string variable
	// and then assign string value into it.
	// var is short for variable
	var cards deck = newDeck()
	hand, remainingCards := cards.deal(5)

	hand.print()
	fmt.Println("Remaining Card below")
	remainingCards.print()
	println(remainingCards.toString())
	remainingCards.saveToFile("test.txt")
	print(newDeckFromFile("test.txt").toString())
}

func newCard() string {
	return "Five of Diamonds"
}
