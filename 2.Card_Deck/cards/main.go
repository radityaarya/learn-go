package main

func main() {
	cards := newDeck()

	takenCard, remainCard := deal(cards, 5)

	takenCard.print()
	remainCard.print()
}
