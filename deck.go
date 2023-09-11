package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

func (d deck) saveToFile(filename string) error {
	// 0666 -> anyone can read and write file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// we will annotate this type string
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log the error and then just quit the program entirely
		log.Println("Error : ", err)
		os.Exit(1)
	}

	s := string(bs) // it is separated by comma
	log.Println("Read Deck from the file : ", s)
	return deck(strings.Split(s, ","))
}
