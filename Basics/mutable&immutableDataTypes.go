package main

import (
	"fmt"
)

// 🏷 Important
// utilizing the pointer replication feature
// no need to return this function is going to actually alter
// the value of the slice passed in as param
func changeFirst(slice []int) {
	slice[0] = 1000
}
func main() {
	// case-1
	// slices are mutable data type
	var x []int = []int{3, 4, 5}
	// y := x will now be pointing to the same slice in memmory any change in x will reflect in y and in y will reflect  in x
	// y here act as another name for x i.e y is new pointer that is pointing to slice {3,4,5} now like x
	y := x
	y[0] = 100
	fmt.Println(x, y)
	x[0] = 3
	// the x slice also get altered even though we changed only y
	fmt.Println(x, y)

	// case-2 (map pointer replication)
	// map are alos mutable data types
	var x2 map[string]int = map[string]int{
		"hello": 3,
	}
	// y2 now points to the map that x2 is pointing to now
	// y2 modification will change x2 & vice-a-versa
	y2 := x2
	y2["y"] = 100
	fmt.Println(x2, y2)

	// case-3
	// arrays (no pointer replication in arrays)
	var x3 [3]int = [3]int{4, 5, 6}

	// no pointer replication in case of arrays
	y3 := x3
	y3[0] = 100
	fmt.Println(x3, y3)

	var x5 []int = []int{3, 5, 6}
	fmt.Println(x5)
	changeFirst(x5)
	fmt.Println(x5)

}
