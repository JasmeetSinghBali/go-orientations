package main

import (
	"fmt"
	"sync"
)

func main() {
	var numMemResources int

	/*acts as memory pool*/
	memPool := &sync.Pool{
		/*every time new instance is creatd it calls this New method that will allocate a mem space to that instance while incrementing the memPool count as well*/
		New: func() interface{} {
			numMemResources++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		/*spawn new go routines*/
		go func() {
			/*allocate mem to eache workers[sub-goroutine] spawned*/
			mem := memPool.Get().(*[]byte)
			fmt.Sprintln("worker using mem for some time...beepbeepboopp")
			/*put mem spaces back to the pool*/
			memPool.Put(mem)
			wg.Done() // to get pass the wg.Wait() line
		}()
	}
	wg.Wait()

	fmt.Printf("%d numMemResources were created", numMemResources)
}
