package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this function gets blocked(not able to iterate further in channel to get other messages) until the message channel does not contain exactly 1 message pushed by the subgoroutine firemissiles
func main() {
	channel := make(chan string)
	// updates how many missiles were fired with a sub goroutine
	go fireMissiles(channel)
	// below commented out results in deadlock as default channel capacity is 0
	// channel <- "yo"
	// iterating through channel updated variable
	// runs from 0 to 3 i.e 4 times and expects 4 messages but only 3 messages are actually their hence deadlock
	// to counter this close the channel in the sub goroutine to make sure the channel is closed
	for message := range channel {
		fmt.Println(message)
	}
	// how the range for approach is working under the hood given below
	// for {
	// 	// get the message & open object from channel
	// 	message, open := <-channel
	// 	// if the channel is no longer open then break the loop
	// 	// as all messages in channel are iterated
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }
}

// this function gets block(not able to iterate further in channel to push message) after 1 message push in channel and becomes unblocked once the main goroutine has accessed the already exisiting message in the channel first
func fireMissiles(channel chan string) {
	rounds := 3
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("you scored", score)
	}
	// close the channel after we are done sending all messages
	close(channel)
}
