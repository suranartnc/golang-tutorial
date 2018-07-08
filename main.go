/*
	Commands
	go run 		= compile, then execute => output
	go build 	= compile => binary file
	go fmt
	go install
	go get
	go test

	2 Types of packages
	1. Executable => Run main.main()
	2. Reusable

	Standard library
	- fmt
	- debug
	- math
	- crypto
	- io
	- encoding
*/

// Package declaration
package main

// Import other packages that we need
import "fmt"

func main() {
	fmt.Println("Hello world!")
}
