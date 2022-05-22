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
	var x string = fmt.Sprintf("%T %v", 10, 20)
	fmt.Println("using Sprintf we can store values", x)
	fmt.Printf("%t \n", bl)
	fmt.Printf("Binary (base 2) representation of Number 4565 is: %b \n", 4565)
	fmt.Printf("Octal (base 8) representation of Number 4565 is: %o \n", 4565)
	fmt.Printf("Decimal (base 10) representation of Number 4565 is: %d\n", 4565)
	fmt.Printf("hexadecimal (base 16) representation of Number 4565 is: %x\n", 4565)
	fmt.Printf("hexadecimal (base 16) representation of Number 4565 in capital letters: %X\n", 4565)
	fmt.Printf("floating points formatter scientific notation Number: %e\n", 232310.92093232121000)
	fmt.Printf("floating points formatter decimal no exponent e  Number: %f \n", 232310.92093232121000)
	fmt.Printf("floating points formatter large exponent Number: %g\n", 232310.920932321210002389283928398293928938928398298327)
	fmt.Printf("Strings formatter-> %s \n", "this is string\n")
	fmt.Printf("Strings formatter-> %q \n", "this is string with double quotes\n")
	fmt.Printf("Width & precision formatter with width/paddding 9 from left and 2 decimal precision -> %9.2f \n", 43.23001)
	fmt.Printf("Width & precision formatter with padding: %-9q yoo this cool \n", "this is left justified by 9")
	fmt.Printf("Width & precision formatter with padding digit 7 as entire length prefixing with 0 from left: \n %07d is \t nice", 45)
}
