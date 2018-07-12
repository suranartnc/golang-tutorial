package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail,com",
			zipCode: 94000,
		},
	}

	// &variable = Give me the memory address of the value this variable is pointing at
	jim.updateName("Jimmy")
	jim.print()
}

// *type = type description of pointer
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointer = Give me the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
