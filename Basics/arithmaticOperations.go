package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 int = 2
	var num2 int = 2
	res1 := num1 + num2
	res2 := num1 - num2
	res3 := num1 / num2
	res4 := num1 * num2
	res5 := num1 % num2

	fmt.Printf("result: %d %d %d %d %d \n", res1, res2, res3, res4, res5)

	var num3 float64 = 8
	var num4 int = 2
	res6 := num3 + float64(num4)

	fmt.Printf("explicit type conversion example \n num1: %f , num2: %d \n result: %f \n", num3, num4, res6)

	var num5 float64 = 0.5
	res7 := float64(math.Cos(num5))

	fmt.Printf("math package cos(0.5) is %f", res7)
}
