package main

import (
	"fmt"
)

type shape interface {
	area() float64 // implemented by circle & rect struct
}

// 2nd interface implemented by circle & rect object
type shape2 interface {
	area() float64
}

// implements the shape interface
type circle struct {
	radius float64
}

// implements the shape interface
type rect struct {
	length  float64
	breadth float64
}

func (c *circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r *rect) area() float64 {
	return r.length * r.breadth
}

// getArea method which takes param s of type shape interface i.e area is accessible here
func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := &circle{4.5}
	r1 := &rect{3, 4}
	fmt.Println("circle and rect..\n", *c1, *r1)
	// shapes is a slice of type interface shape that can only access area method
	shapes := []shape{c1, r1}
	fmt.Println("Shapes slice of type interface shape:/n", shapes)
	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
