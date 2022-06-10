package main

import "fmt"

func main() {
	ans := 10

	switch ans {
	case 1:
		fmt.Println("Its fine!!")
	case -1, 10:
		fmt.Println("its a dual case check intresting!!")
	default:
		fmt.Println("daddy chill")
	}

	switch {
	case ans > 5:
		fmt.Println("nice style for switch statement here")
	}

}
