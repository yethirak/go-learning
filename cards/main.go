package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()
	// cards := deck{newCard(), "Ace of Spade"}
	// cards = append(cards, "Six of Spades")
	cards := newDeck()
	// cards := newDeckFromFile("my_cards1")
	cards.shuffle()
	// fmt.Println(cards)
	// cards.printDeck()
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Hand: " + hand.toString())
	// hand.printDeck()
	fmt.Println("Remaining Deck: " + remainingDeck.toString())
	// remainingDeck.printDeck()
	// cards.saveToFile("my_cards")
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
