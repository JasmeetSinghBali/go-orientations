package main

import "fmt"

func main() {
	// will overflow as uint8 cannot denote 260
	// var number uint8 = 260
	var number uint = 100023345
	var number2 uint16 = 260
	fmt.Println("Hello go!", number, number2)
}
