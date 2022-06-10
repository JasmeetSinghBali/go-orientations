package main

import (
	"fmt"
)

// CALL BY REFERENCE (passing the actual address where the string is stored in memory RAM)
// str *string will give us a pointer here in param
func changeValue(str *string) {
	// deref and update the value via str pointer
	*str = "changedCallByRef"
}

// CALL BY VALUE (passsing just the copy)
func changeValue2(str string) {
	str = "changedCallByValueWONTWORK"
}

func main() {
	x := 7 // 7 is the value inside x
	// to look at the refer of the variable i.e address it is stored at
	fmt.Println(&x) // addrress where x is stored i.e ref or pointer of x
	y := &x         // y is now pointing to location of x
	fmt.Println(x, y)
	*y = 8 // deref is used to access the address to which the pointer is pointing to since y is points to address of x *y will help us to deref and access/update the value block at this address that is the address of x

	fmt.Println(x, y)

	tochange := "hi"
	// &tochange passs pointer to string
	fmt.Println(tochange)
	changeValue(&tochange)
	fmt.Println(tochange)
	/// no effect call by value on the actual address
	changeValue2(tochange)
	fmt.Println(tochange)

	toChange := "Hello POinter"
	var pointer *string = &toChange
	// address where the pointer point will be stored in pointer
	fmt.Println(pointer)
	// value of the location whose address pointer is holding
	fmt.Println(*pointer)
}
