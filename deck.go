package main

import "strings"

// Create a New type of deck
// which is a slices of string.
type deck []string

/*
 * Return the list of cards are consist of string
 */
func newDeck() deck {
	cards := deck{}

	// we're not making this a deck because a deck is really supposed to be used with actual playing cards.
	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}

	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
 * Return a string
 * ["red", "yellow", "blue"] -> condense it down to a string of red, comma, yellow, comma, blue -> "red,yellow,blue"
 */
func (d deck) toString() string {
	return strings.Join(d, ",")
}
