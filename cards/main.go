package main

func main() {
	cards := newDeckFromFile("myCards.txt")
	hand, remainingCards := deal(cards.shuffle(), 3)
	hand.print()
	remainingCards.print()

	cards.saveToFile("myCards.txt")
}
