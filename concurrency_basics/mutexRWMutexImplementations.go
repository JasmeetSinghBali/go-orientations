package main

import (
	"fmt"
	"sync"
	"time"
)

// package level global variable
var (
	// lock act as mutex controller variable
	lock sync.Mutex
	// read write mutex
	rwlock sync.RWMutex
	count  int
)

func main() {
	//basics()
	readAndWrite()
}

// basics mutex goroutine
func basics() {
	itrns := 1000000
	// spawning sub goroutines
	for i := 0; i < itrns; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Resulted Count is:", count)
}

func increment() {
	lock.Lock() // now only one sub goroutine has the actual access to the count variable.
	count++
	// count++ under the hood
	// temp := count
	// temp = temp+1
	// count = temp
	// so the increment is not atomic so for some subroutines it will miss the increments if we dont use lock & unlock mutex flows
	lock.Unlock()
}

// read/write mutex goroutine
func readAndWrite() {
	go read()  // subgoroutine for reading
	go read()  // another sybgoroutine for reading
	go write() // subgoroutine for writing

	time.Sleep(5 * time.Second)
	fmt.Println("Done...")
}
func read() {
	rwlock.RLock()         // apply read lock
	defer rwlock.RUnlock() // unlock just before finally returning from read() subgoroutine

	fmt.Println("Reading locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Reading unlocking")
}

func write() {
	rwlock.Lock()         // apply full lock
	defer rwlock.Unlock() // unlock just before finally returning from write() subgoroutine

	fmt.Println("full write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("full write lock unlocking")
}
