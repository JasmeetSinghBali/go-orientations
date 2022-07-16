package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/*Adding 1 to sum variable via sync/atomic*/
	var sum int64
	fmt.Println(sum)
	atomic.AddInt64(&sum, 1)
	fmt.Println(sum)

	/*Normal way of adding 1 to sum variable via mutex*/
	var mu sync.Mutex
	mu.Lock()
	sum = sum + 1
	mu.Unlock()
	fmt.Println(sum)

	var diffSum int64
	fmt.Println(atomic.LoadInt64(&diffSum))
	atomic.StoreInt64(&diffSum, 1)
	fmt.Println(diffSum)

	/*🎯IMP: Perform atomic concurrent operations on any customized data/object type*/
	var av atomic.Value
	john := myCustomType{name: "john"}
	av.Store(john)
	fmt.Println(av.Load().(myCustomType).name)

	// waitgroup to prevent any unwanted race conditions
	var wg sync.WaitGroup
	wg.Add(1) // since we are only spawing one go sub routine for modification of the av atomic.Value

	// go routine that make modification to the av atomic value
	go func() {
		/*get value in av with typecasting it to myCustomType struct*/
		j := av.Load().(myCustomType)
		// update the content
		j.name = "Not john"
		// store the updated one in the av var
		av.Store(j)
		j.name = "john again" // 🎯:IMP this wont have any effect as av.Store(j) we are using pass by value not pass by reff refer: Store() method cntrl+click it is interface{} type and not pointer , in contrast to this the inbuild Addint64 and so on pass by reff i.e pointer ref
		// signal the waitgroup
		wg.Done()
	}()
	// wait for all goroutines to finish to avoid race conditions
	wg.Wait()
	fmt.Println(av.Load().(myCustomType).name)
}

type myCustomType struct {
	name string
}
