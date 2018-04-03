package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// create new deck from a file
func newDeckFromFile(filename string) deck {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(file), ",")
	return s
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

// to save deck (string) to a file
func (d deck) saveToFile(fileName string) error {
	deckInByte := []byte(d.arrToString())
	return ioutil.WriteFile(fileName, deckInByte, 0666)
}

// to shuffle card's order on deck
func (d deck) shuffle() deck {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// check '~/Algorithms/randomNumber.go'

	// shuffle time!
	for i := range d {
		newPosition := r.Intn(len(d))

		// how to swap value of array on GO 😎
		d[i], d[newPosition] = d[newPosition], d[i]
	}

	return d
}
