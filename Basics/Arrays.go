package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 10
	fmt.Println(arr[0])
	arr2 := [3]int{45, 100, 1}
	fmt.Println(len(arr2))
	sum := 0
	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}
	fmt.Println(sum)
	// multidimensional array
	// [lengthOfArray][interiorElemInside1dArray]
	arr3 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr3[0][1], arr3)
}
