package main

import (
	"fmt"
	"net/http"
	"time"
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

	// Without channel, main thread finishes before child threads
	c := make(chan string)

	for _, link := range links {
		// Create a new thread to run this function
		go checkLink(link, c)
	}

	for l := range c {
		// Pause main thread for 5 seconds
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
	}
}

/*
	Channels
	- channel <- data					= send data to channel
	- variable <- channel 		= wait for data from channel, then assign it to a variable
	- fmt.Println(<-channel)	= wait for data from channel, then log it out
*/

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
