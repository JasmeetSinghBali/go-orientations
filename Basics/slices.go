package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original array\n", arr)
	var s []int = arr[1:3]
	// fmt.Println("copy the entire array into slice with [:] \n", s)
	fmt.Println("from 1 to 3 but do not include 3rd index [1:3] \n", s)
	fmt.Println("capacity of the slice: with cap() \n", cap(s))
	// note the capcaity is the number of elements after the initial size index
	// [1,2,3,4,5]
	// since the slice started from index 1 we have 4 elem 2,3,4,5 so 4 space and that becomes the capacity of the new slice
	fmt.Println("extend the capacity of the slice with the slice capacity", s[:cap(s)])

	// make a slice
	var a []int = []int{5, 6, 7, 8, 9}
	fmt.Println("making a a slice with initial values\n", a)
	fmt.Println("capacity of the slice of slice a[:3]\n", cap(a[:3]))

	// we cant increase the size of slice, instead create a new slice and append old values with addon size
	// append(sliceasFirstArgument,argumentToAddToSlice)
	fmt.Println("append 10 with [5,6,7,8,9] slice: \n", append(a, 10))

	// make a slice with keyword make
	makedSlice := make([]int, 5)
	fmt.Printf("This is a slice Type: %T", makedSlice)
}
