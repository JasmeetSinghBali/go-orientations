package main

import (
	"fmt"
)

// custom struct type
// a collection of different x,y pairs can exist now with different values
type Point struct {
	x int32
	y int32
}

func changeXPassByRef(pt *Point) {
	pt.y = 100
}

func changeXPassByValue(pt Point) {
	pt.y = 100
}

// custom embeded struct with a pointer to Point type declared above
type Circle struct {
	radius float64
	center *Point
}

func main() {
	var p1 Point = Point{1, 2}
	// p2 is type of Point just like p1 just d/f syntax
	p2 := Point{-5, -7}
	fmt.Println("This is custom type struct with x & y as int32\n", p1, p2)
	fmt.Println("access the y of p2->\n", p2.y)
	// default value will be set a/c to type if u dont pass it
	p3 := Point{y: 4}
	fmt.Println("Only set value of y\n", p3)
	point := &Point{x: 3}
	fmt.Println("This is a pointer to the x,y custom struct with address pointing to with value->\n", point, &point, *point)
	// pass the pointer i.e ref of Point type {x:3,y:0}
	changeXPassByRef(point)
	fmt.Println("After change in x with pointer ref/n", *point)
	point2 := Point{x: 3}
	fmt.Println("CASE:2 \npass by value no address or pointer use no change as only a copy is passed\n", point2)
	changeXPassByValue(point2)
	fmt.Println(point2)
	// the structs dont require deref operator to actually access the values
	point.x = 100
	fmt.Println("No need of deref * in struct, directly change value from {x:3 y:100} to \n", point)

	c1 := Circle{3.46, &Point{y: 2}}
	fmt.Println("This is embedded pointer\n", c1)
	fmt.Println("with the values of center: \n", c1.center)
	fmt.Println("with the values of center the x value: \n", c1.center.x)
	c1.center.x = 100
	fmt.Println("changed the values of center: \n", c1.center)
}
