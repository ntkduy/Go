package main

import "fmt"

func main() {
	/* Variable Declaration */
	card := "Ace of Spades"
	card = "Five of Diamonds"

	card = newCard()
	fmt.Println(card)
	printState()

	/* Slices */
	cards := []string{newCard(), "Ace of Spades"}
	cards = append(cards, "Six of Spades")

	/* For Loop */
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	/* OO Approach vs GO Approach */
	// deckCards := deck{newCard(), "Ace of Spades"}
	// for i, deckCard := range deckCards {
	// 	fmt.Println(i+2, deckCard)
	// }
	// This code line behave exactly the same with the for loop above
	// deckCards.print()

	// Using newCard func in deck.go
	newCardDeck := newDeck()
	// newCardDeck.print()

	// Using deal func in deck.go
	hand, remainingCards := deal(newCardDeck, 5)
	hand.print()
	remainingCards.print()

}

func newCard() string {
	return "string from the newCard() function"
}
