package main

import "fmt"

func main() {
	// x := 5
	// y := 6.5
	x := "tim"
	y := "Tim"
	val := x == y
	fmt.Printf("%t \n", val)
	fmt.Printf("Ascii comparison between a < b:  %t\n", 'a' < 'b')
	fmt.Printf("Ascii comparison between A(65) > b:  %t", 'A' > 'b')
}
