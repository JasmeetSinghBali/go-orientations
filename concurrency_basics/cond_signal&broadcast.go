package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	gettingReadyForMissionWithCond()
	broadcastStartOfMission()
}

var ready bool

/*Signalling without cond [THE WRONG WAY]*/
// this approach has to go through a large number of work intervals to determine that worker is ready for mission
func gettingReadyForMission() {
	/*1st goroutine gettingReady() go routine that updates the ready var*/
	go gettingReady()
	/*2nd go routine that keeps checking wheather worker is ready*/
	workIntervals := 0
	for !ready {
		/*spending 5 seconds to check for wheather the worker is ready as without time.sleep() a lot of workIntervals have to waited on to check wheather the worker is ready*/
		//time.Sleep(5 * time.Second)
		workIntervals++
	}
	fmt.Printf("Workers are now ready, after %d work intervals.\n", workIntervals)
}

/*📝Signalling via cond [THE RIGHT WAY]*/
// the go routine code that checks ready dont have to wait for long or check large number of workIntervals as the gettingReadyWithCond() is going to signal when the ready variable is set to true
func gettingReadyForMissionWithCond() {
	/*condition variable*/
	cond := sync.NewCond(&sync.Mutex{})
	/*in params no need to specify &cond as the NewCond automatically returns the address and stores in cond*/
	/*passing cond address of the Cond variable to the gettingReadyWithCond go sub-routine*/
	go gettingReadyWithCond(cond)
	workIntervals := 0

	// 📝locking the cond variable
	cond.L.Lock()
	for !ready {
		workIntervals++
		// 📝waiting for other go sub-routines to finish the task of getting ready
		cond.Wait()
	}
	// 📝unlocking the cond variable
	cond.L.Unlock()

	fmt.Printf("Workers are now ready, after %d work intervals.\n", workIntervals)
}

/*📝 broadcast mission with waitgroup and broadcast pair*/
func broadcastStartOfMission() {
	beeper := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3) // as we have 3 sub goroutines call execution to wait for
	/*📝 each beeper holds the NewCond address that is passed to each of the go subroutine call standByForMission()*/
	standByForMission(func() {
		fmt.Println("Minion 1 started mission")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Minion 2 started mission")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Minion 3 started mission")
		wg.Done()
	}, beeper)
	// 📝 broadcasting that all gosub routines have completed executing
	beeper.Broadcast()
	wg.Wait() // wait for all sub go routines to finish their task
	fmt.Println("All minions have started the missions")
}

/*📝 standByForMission() call locks/unlocks the cond*/
func standByForMission(fn func(), broadcaster *sync.Cond) {
	var wg sync.WaitGroup
	wg.Add(1) // as we are waiting for only 1 sub goroutine
	go func() {
		wg.Done()
		broadcaster.L.Lock()
		defer broadcaster.L.Unlock()
		broadcaster.Wait()
		fn()
	}()
	wg.Wait()
}

func gettingReady() {
	sleep()
	ready = true
}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal() // 📝 signalling the gettingReadyForMissionWithCond() goroutine that the ready variable is set to true.
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
