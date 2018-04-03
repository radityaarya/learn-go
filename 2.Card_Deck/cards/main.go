package main

import "fmt"

func main() {
	cards := newDeck()

	cards.shuffle()

	takenCard, remainCard := deal(cards, 5)

	fmt.Println("TAKEN CARDS :", takenCard.arrToString()+"\n\n"+"REMAIN CARDS :", remainCard.arrToString())
	cards.saveToFile("my_cards")

}
