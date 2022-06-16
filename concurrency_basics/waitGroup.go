package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitgrp sync.WaitGroup
	soldList := []string{"ronny", "kapil", "dhigman"}
	// how many go routines we want to move in wait state
	waitgrp.Add(len(soldList))
	for _, soldier := range soldList {
		// pass by ref wait grp pointer to attackSold
		go attackSold(soldier, &waitgrp)
	}
	// waiting until the attackSold is done with all slice string values of soldiers
	waitgrp.Wait()
	fmt.Println("attack was a success")
}

func attackSold(target string, waitgrp *sync.WaitGroup) {
	fmt.Printf("Attacked soldier: %s\n", target)
	waitgrp.Done()
}
