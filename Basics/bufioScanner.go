package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// step-1 variable scanner is of type buffer input output create a NewScanner class object with oper systems input with Stdin
	scanner := bufio.NewScanner(os.Stdin)
	// to have a question text
	fmt.Printf("Type your year of birth: ")
	// step-2 now this scanner variable will scan the input by user in console
	scanner.Scan()
	// step-3 store the scanned input in the input variable
	// _ or err refers to the error if incase the input from user is not converted into number for some reason
	// use _ if u dont care about error
	// note input will be 0 if type conversion fails
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Your will be %d year's old by end of 2022!", 2022-input)
}
