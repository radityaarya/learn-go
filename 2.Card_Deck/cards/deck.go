package main

import "fmt"

type deck []string

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

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index+1, card)
	}
}
