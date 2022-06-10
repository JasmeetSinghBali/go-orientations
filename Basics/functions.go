package main

import "fmt"

func test(x, y int) (z1, z2 int) {
	// important defer runs just before return statement is read
	defer fmt.Printf("[defer] This will happen just before the function returns can be used for \n 1. cleaning code \n2. doing some specific job just before function returns\n")
	z1 = x + y
	z2 = x - y
	fmt.Printf("not yet reached return of the fucntion\n")
	// automatically returns the z1 and z2 these are labelled returns from function
	return
}

func main() {
	res1, res2 := test(1, 2)
	fmt.Println(res1, res2)
}
