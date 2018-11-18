package main

func main() {
	// cards := newDeck()

	// // receiver on function
	// // cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.saveToFile("MyCards")

	cards := newDeckFromFile("My")
	cards.print()
}
