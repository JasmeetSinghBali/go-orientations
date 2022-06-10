package main

import "fmt"

func main() {
	var a []int = []int{2, 34, 56, 69, 11, 4, 5, 2, 34}
	fmt.Println(a)
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// RANGE & SLICES
	// _ replaces the variable we dont care about so that compiler dont scream on us
	// note _ value cannot be accessed here
	// for _, element := range a {
	// 	fmt.Printf("%d\n", element)
	// }

	// find duplicate elements in the slice
	for i, elm := range a {
		for j := i + 1; j < len(a); j++ {
			if elm == a[j] {
				fmt.Println(elm)
			}
		}
	}
}
