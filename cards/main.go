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
*/

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
