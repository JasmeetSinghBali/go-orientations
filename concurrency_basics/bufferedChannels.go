package main

import (
	"fmt"
)

func main() {
	// note- default capacity of channel is 0
	// buffered-channel i.e channel with customized capacity
	// capacity - 1 given to channel to be able to update & access the variable in channel by main
	channel := make(chan string, 2)
	channel <- "first mesg"
	channel <- "second mesg"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
