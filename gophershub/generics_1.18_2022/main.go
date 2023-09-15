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

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

// 💫 a generic struct
type User[T CustomData] struct {
	ID   int
	Name string
	Data T // this data cud be anything
}

// 💫 a generic map
// here the key can be generic or maybe value
// NOTE- The valid key for a map can only have comparable types
// example a(int) == b(int) , a(string) == b(string), *a == *b i.e any value that can be compared with the same value of its type
type CustomMap[T comparable, V int | string] map[T]V

func main() {
	resultInt := MapItOver([]int{1, 2, 3}, func(n int) int {
		return n * 7
	})
	fmt.Println(resultInt)
	resultFloat := MapItOver([]float64{1.2, 2.3, 3.1}, func(n float64) float64 {
		return n * 7
	})
	fmt.Println(resultFloat)

	userWithInt := User[int]{
		ID:   1,
		Name: "JohnDoe",
		Data: 3,
	}
	fmt.Println(userWithInt)
	userWithString := User[string]{
		ID:   1,
		Name: "JohnDoe",
		Data: "somerandomfacts!",
	}
	fmt.Println(userWithString)

	mapWithIntKey := make(CustomMap[int, string])
	mapWithIntKey[3] = "yo"
	fmt.Println(mapWithIntKey)
	mapWithStringKey := make(CustomMap[string, int])
	mapWithStringKey["yo"] = 3
	fmt.Println(mapWithStringKey)
	mapWithStringKeyValue := make(CustomMap[string, string])
	mapWithStringKeyValue["yo"] = "yo"
	fmt.Println(mapWithStringKeyValue)
}
