package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Diamonds"

	card = newCard()
	fmt.Println(card)
	printState()
}

func newCard() string {
	return "string from the newCard() function"
}
