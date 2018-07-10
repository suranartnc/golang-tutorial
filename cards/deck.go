package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

/*
	Receiver Function
	- cards = actual value of deck
	- deck 	= only this type can use the method
*/
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
