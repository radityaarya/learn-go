package main

func main() {
	cards := newDeck()

	// takenCard, remainCard := deal(cards, 5)

	// takenCard.print()
	// remainCard.print()
	// fmt.Println(cards.arrToString())
	cards.saveToFile("my_cards")
}
