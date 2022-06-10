package main

import (
	"fmt"
)

func returnsFunction(x string) func() {
	// param of outside function accessible to the nested function inside of it
	return func() { fmt.Println(x) }
}

func main() {
	returnsFunction("hello")()
	returnsFunction("nice function closure u got their")()
}
