package main

import "fmt"

func main() {

	x := "tiM"
	if x == "tim" {
		fmt.Printf("tim is a good boy!")

	} else if x == "Tim" {
		fmt.Printf("tim is a bad boy!")
	} else {
		fmt.Printf("No idea about tim character! I am not the judge")
	}
}
