package main

import (
	"fmt"
	"strings"
)

type deck []string

// create and return a list of cards (array of strings)
func newDeck() deck {
	cards := deck{}

	cardSymbols := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, symbols := range cardSymbols {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+symbols)
		}
	}

	return cards
}

// print out the content of a deck of cards
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index+1, card)
	}
}

// split deck into 2 => hand and deck. total cards on hand called DEAL
func deal(d deck, sizeOfDeal int) (deck, deck) {
	cardsOnHand := d[0:sizeOfDeal] // shorter syntax => d[:sizeOfDeal] only if start from 0
	cardsOnDeck := d[sizeOfDeal:]

	return cardsOnHand, cardsOnDeck
}

// convert array to string to string comma-separated
func (d deck) arrToString() string {
	return strings.Join(d, ",")
}
