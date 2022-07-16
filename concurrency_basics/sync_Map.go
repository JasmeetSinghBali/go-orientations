package main

import (
	"fmt"
	"sync"
)

func main() {
	regularMap := make(map[int]interface{})
	syncMap := sync.Map{}

	//put
	regularMap[0] = 0
	regularMap[1] = 1
	regularMap[2] = 2

	syncMap.Store(0, 0)
	syncMap.Store(1, 1)
	syncMap.Store(2, 2)

	// get
	regularValue, regularOk := regularMap[0]
	fmt.Println(regularValue, regularOk)

	syncValue, syncMapOk := syncMap.Load(0)
	fmt.Println(syncValue, syncMapOk)

	// delete
	regularMap[1] = nil
	syncMap.Delete(1)

	/*📝LoadAndDelete will load i.e store the value that is being deleted, its equivalen to below 2 line code snippet*/
	// i.e deletedValue = regularMap[2]
	// delete(regularMap,2)
	syncValue, loaded := syncMap.LoadAndDelete(2)
	/*📝LoadAndDelete inner working*/
	mu := sync.Mutex{}
	mu.Lock()
	deletedValue := regularMap[2]
	delete(regularMap, 2)
	mu.Unlock()
	fmt.Println(syncValue, loaded, deletedValue)

	// get and put
	// syncValue, loaded = syncMap.LoadOrStore(1,1)
	// mu := sync.Mutex{}
	// mu.Lock()
	// storedValue, storedValueOk = regularMap[1]
	// if regularOk{
	// 	regularMap[1] = 1
	// 	regularValue = regularMap[1]
	// }
	// mu.Unlock()
	// fmt.Println(syncValue,storedValue)

	// range
	for key, value := range regularMap {
		fmt.Print(key, value, " | ")
	}

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Print(key, value, " | ")
		return true
	})
}
