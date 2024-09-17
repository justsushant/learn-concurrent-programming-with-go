package main

import (
	"fmt"
	"time"
)

// find race condition

func addNextNumber(nextNum *[101]int) {
	i := 0
	for nextNum[i] != 0 { i++ }
	nextNum[i] = nextNum[i-1] + 1
}

func main() {
	nextNum := [101]int{1}
	for i := 0; i < 100; i++ {
		// multiple goroutines are trying to 
		// modify elements from the same array 
		// since pointer of the array was passed to all goroutines
		go addNextNumber(&nextNum)
	}
 	
	for nextNum[100] == 0 {
		println("Waiting for goroutines to complete")
		time.Sleep(10 * time.Millisecond)
 	}
 	
	fmt.Println(nextNum)
}