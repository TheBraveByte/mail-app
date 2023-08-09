package main

import (
	"fmt"
	"time"
)

func main() {
	// Declare an unbuffered channel of integers.
	intChan := make(chan int)
	stringChan := make(chan string)

	// Send an integer to the channel.
	go func() {
		intChan <- 299
	}()

	// Send a string to the channel.
	go func() {
		stringChan <- "Hello, Gophers."
	}()

	// Receive from the channels in a separate goroutine.
	go func() {
		select {
		case i := <-intChan:
			fmt.Println("Received an integer:", i)
		case s := <-stringChan:
			fmt.Println("Received a string:", s)
		}
	}()

	// Wait for 1 second before ending the program.
	time.Sleep(1 * time.Second)
}
