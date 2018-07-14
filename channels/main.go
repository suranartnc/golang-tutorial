package main

import (
	"fmt"
	"net/http"
)

/*
	Go Routines (Threads in Go)
	- Main routine created when we launched our program
	- Child routine created by the "go" keyword

	One vs Multiple CPU Core
	- One CPU Core - Run only one routine at a time until it finishes or it makes a blocking call
	- Multiple CPU Core - Run one routine on each logical core

	Concurrency
	- If one thread blocks, another one is picked up

	Parallelism
	- Use multiple CPU Core to execute multiple threads at the same time
*/

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		// Create a new thread to run this function
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
