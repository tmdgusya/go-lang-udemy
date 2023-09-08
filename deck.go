package main

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
