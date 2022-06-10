package main

import (
	"fmt"
)

func test() int {
	fmt.Println("Noice")
	return 1
}
func test3(myfunc func(int) int) {
	fmt.Println("This is a function with param as another function...\n")
	fmt.Println(myfunc(7))
}

func main() {
	// passing refference fo function to variable
	x := test
	x()
	// nested function & inline function
	test2 := func(x int) (z1 int) {
		fmt.Println("this is a nested function which is refferenced to a test variable")
		z1 = x * x
		return
	}
	// passing function as params
	test3(test2)
}
