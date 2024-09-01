package main

import (
	"fmt"
	"time"
)

func receivingFromMultipleChannels() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func(channel chan<- string) {
		time.Sleep(1 * time.Second)
		chan1 <- "from channel 1"
	}(chan1)

	go func(channel chan<- string) {
		time.Sleep(2 * time.Second)
		chan1 <- "from channel 2"
	}(chan2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received", msg2)
		}
	}
}

func withTimeOutFromMultipleChannels() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func(channel chan<- string) {
		time.Sleep(2 * time.Second)
		chan1 <- "from channel 1"
	}(chan1)

	go func(channel chan<- string) {
		time.Sleep(4 * time.Second)
		chan1 <- "from channel 2"
	}(chan2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
			return
		}
	}
}
func main() {
	receivingFromMultipleChannels()
	withTimeOutFromMultipleChannels()
}
