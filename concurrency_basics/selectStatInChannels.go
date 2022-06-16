package main

import "fmt"

func main() {
	chn1, chn2 := make(chan string), make(chan string)

	go electCaptain(chn1, "soldier-1 became captain")
	go electCaptain(chn2, "soldier-2 became captain")

	// it will randomly select a case & then execute that case
	select {
	case message := <-chn1:
		fmt.Println(message)
	case message := <-chn2:
		fmt.Println(message)
	}

	electCaptainfairly()
}

func electCaptain(captan chan string, message string) {
	captan <- message
}

func electCaptainfairly() {
	chn1 := make(chan string)
	close(chn1)
	chn2 := make(chan string)
	close(chn2)

	var sol1count, sol2count int
	for i := 0; i <= 1000; i++ {
		select {
		case <-chn1:
			sol1count++
		case <-chn2:
			sol2count++
		}
	}
	fmt.Printf("soldier-1Count: %d soldier-2Count: %d\n", sol1count, sol2count)
}
