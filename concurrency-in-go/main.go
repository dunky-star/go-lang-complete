package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello! ", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // Simulating slow greeting
	fmt.Println("Hello! ", phrase)
	doneChan <- true
}

func main() {
	// To acivate goroutine. Add "go" before the function call
	// To run the functions in a non blocking way
	go greet("Nice to meet yah!")
	go greet("How are you?")
	done := make(chan bool)
	go slowGreet("How ... are ... you ... ?", done)
	go greet("I hope you're liking it here in North America")
	<-done
}
