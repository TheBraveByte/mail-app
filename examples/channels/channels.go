package main

import (
"fmt"
"time"
)


func main() {
// Declare a buffered channel of integers with a capacity of 2.
bufferedChan := make(chan int, 2)

// Unbuffered channel with no capacity
unbufferedChan := make(chan string)

// Send two integers to the channel.
bufferedChan <- 10
bufferedChan <- 20

// Start a goroutine to receive data from the unbuffered channel.
go func() {
data := <-unbufferedChan
fmt.Println("Received from unbuffered channel:", data)
}()

// Wait for a short period before sending and receiving more data.

time.Sleep(time.Second)
// Send some data to the unbuffered channel.
unbufferedChan <- "Hello, Gophers."

// Receive some data from the buffered channel.
data1 := <-bufferedChan
data2 := <-bufferedChan

// Print out the received data.
fmt.Println("Received from buffered channel:", data1, data2)

}
