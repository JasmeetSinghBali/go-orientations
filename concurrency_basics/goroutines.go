package main

import (
	"fmt"
	"time"
)

// idea-> to attack all the soldiers at one time concurrently so that the execution time can be reduced
func main() {
	start := time.Now()
	// defer function executes right before the main function is going to return
	// shows the time taken to perform attack on all soldiers
	defer func() {
		fmt.Println(time.Since(start))
	}()
	fmt.Println("Concurrrency via Goroutines")

	soldiersList := []string{"kapil", "dhiman", "yadav"}
	for _, sold := range soldiersList {
		//attack(sold)
		// new process spawn
		// it will now only take 1 second no matter how many soldiers are their
		go attack(sold)
	}
	// to sleep for 2 seconds so that the attack on all soldiers can take place
	time.Sleep(time.Second * 2)

}

func attack(target string) {
	fmt.Println("Throwing missiles at", target)
	time.Sleep(time.Second)
}
