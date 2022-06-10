package main

import (
	"fmt"
)

func main() {
	fmt.Printf("style1\n")
	x := 1
	for x <= 10 {
		fmt.Println(x)
		x += 2
	}
	fmt.Printf("style2\n")
	for y := 10; y <= 22; y += 3 {
		if y == 19 {
			break
		}
		if y == 13 {
			continue
		}
		fmt.Println(y)
	}
	fmt.Printf("just fun\n")
	for z := 0; z <= 1000; z++ {
		if z != 0 && z%3 == 0 && z%5 == 0 && z%7 == 0 {
			fmt.Println(z)
		}
	}

}
