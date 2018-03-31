package main

func main() {
	cards := newDeckFromFile("my_csards")

	// takenCard, remainCard := deal(cards, 5)

	// takenCard.print()
	// remainCard.print()
	// fmt.Println(cards.arrToString())
	// cards.saveToFile("my_cards")
	cards.print()
}
