package main

import (
	"fmt"
	"time"
)

/*
  Note: this program has a race condition for demonstration purposes
  You might need to run this multiple times to trigger the race condition

  There is a race condition because of the following
  1. There is a fininte number of goroutines running, all of which are reading from and writing to a shared array without synchronization, therefore it is not guaranteed that each goroutine will write to a distinct index i.
    - Because they aren't synchronized, multiple threads could end up stopping their for loops on the same i, and therefore write to the same i in the array.
	- it is possible that all threads will have finished writing to an index i prior to index 100 being reached, so the for loop in the main thread of executino will never terminate
*/

func addNextNumber(nextNum *[101]int) {
	i := 0
	for nextNum[i] != 0 {
		i++
	}
	nextNum[i] = nextNum[i-1] + 1
}

func main() {
	nextNum := [101]int{1}
	for i := 0; i < 100; i++ {
		go addNextNumber(&nextNum)
	}
	for nextNum[100] == 0 {
		fmt.Println("Waiting for goroutines to complete")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(nextNum)
}