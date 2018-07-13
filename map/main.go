package main

import (
	"fmt"
)

func main() {
	// Create an empty map

	// Equivalent to
	// var colors map[string]string
	colors := make(map[string]string)

	colors["red"] = "#ff000"
	colors["green"] = "#4bf745"

	delete(colors, "red")

	fmt.Println(colors)
}
