package main

import "fmt"

func main() {
	// will overflow as uint8 cannot denote 260
	// var number uint8 = 260
	// explicit declaration
	var number uint = 100023345
	var number2 uint16 = 260
	// implicit declaration
	var number3 = "yo this is string"
	var number4 = 24.4523
	number5 := -34812
	number6 := 345
	var bl bool
	fmt.Println("Hello go!", number, number2, number3)
	// %T gives us the type of variable
	fmt.Printf("%T \n", number3)
	fmt.Printf("%T\n", number4)
	fmt.Printf("%T\n", number5)
	fmt.Printf("%T\n", number6)
	fmt.Println("[default value of uninitialized boolean variable is ] -> ", bl)
}
