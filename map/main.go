package main

import (
	"fmt"
)

/*
	Map
	- represent a collection of related properties
	- reference type
	- all keys must be same type (all string/all int)
	- all values must be the same type

	Struct
	- represent a thing with a lot of different properties
	- value type
	- all keys must be only string
	- value can be of different typess
*/

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	// for key, value := range map
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
