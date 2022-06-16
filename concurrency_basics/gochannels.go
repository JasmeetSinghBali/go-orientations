package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	// channel
	targetSignal := make(chan bool)
	sold := "kapil"
	go attack(sold, targetSignal)

	// IMPORTANT Case: DEADLOCK main also trying to change bool channel variable
	// results in deadlock as main and spawned gorutine are trying to access the same variable in channel
	// deadlock is resolved via the golang buffered channels
	targetSignal <- false

	// receive value from channel <-channelName
	fmt.Println(<-targetSignal)
}

func attack(target string, signal chan bool) {
	fmt.Println("targetting soldier:", target)
	// <-use this in channels to pass a value from a channel to refference
	// sending true to channel
	signal <- true
}
