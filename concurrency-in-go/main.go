package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneCahn chan bool) {
	fmt.Println("Hello! ", phrase)
	doneCahn <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // Simulating slow greeting
	fmt.Println("Hello! ", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// To acivate goroutine. Add "go" before the function call
	// To run the functions in a non blocking way
	done := make(chan bool)
	//dones := make([]chan bool, 4)
	go greet("Nice to meet yah!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ... ?", done)
	go greet("I hope you're liking it here in North America", done)

	for range done {
		<-done
	}

}
