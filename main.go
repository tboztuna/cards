package main

func main() {

	cards := newDeck()
	cards.saveToFile("my_file.txt")
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
}
