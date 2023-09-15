package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 💫 a generic mapping function
// input: [1,2,3],(n) => n*2
// output: [2,4,6]
// NOTE- constraints.Ordered contain all the int & string types
func MapItOver[T constraints.Ordered](val []T, processFunc func(T) T) []T {
	var newValues []T
	for _, v := range val {
		newValue := processFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	resultInt := MapItOver([]int{1, 2, 3}, func(n int) int {
		return n * 7
	})
	fmt.Println(resultInt)
	resultFloat := MapItOver([]float64{1.2, 2.3, 3.1}, func(n float64) float64 {
		return n * 7
	})
	fmt.Println(resultFloat)
}
