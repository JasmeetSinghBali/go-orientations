package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mangoesFound bool

func main() {
	/*create a waitgroup*/
	var wg sync.WaitGroup

	/*sending 100 sibling*/
	wg.Add(100)

	/*📝 once that will make sure the first sibling that finds mangoes inform others so that other can take a chill pill*/
	var once sync.Once

	/*100 siblings trying to buy mangoes*/
	for i := 0; i < 100; i++ {
		/*go subroutine for dispatched sibling*/
		go func() {
			if foundMangoes() {
				/*📝once.Do wrapper makes sure the function reff passed inside it is executed exactly once*/
				once.Do(signalMangoesFound)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	checkMangoesFound()
}

func checkMangoesFound() {
	if mangoesFound {
		fmt.Println("Mission Success \nMangoes found,all siblings now returned home, no more siblings dispatched now... ")
	} else {
		fmt.Println("Mission failed \nAll sibling shud go home now, while sibling at home shud stay at home...")
	}
}

func signalMangoesFound() {
	mangoesFound = true
	fmt.Println("Signalling we found good mangoes...")
}

func foundMangoes() bool {
	rand.Seed(time.Now().UnixNano())
	// 1 in 10 siblings able to find the mangoes
	return 0 == rand.Intn(10)
}
