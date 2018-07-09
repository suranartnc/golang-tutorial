package main

import (
	"fmt"
)

/*
	Basic Go Types
	- bool
	- string
	- int
	- float64

	Array vs Slice
	- Array = List that have fixed length
	- Slice = List that can grow or shrink
*/

func main() {
	// Create a Slice
	cards := []string{"Ace of Diamonds", newCard()}

	// append() is immutable
	cards = append(cards, "Six of Spades")

	// For Loops
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
