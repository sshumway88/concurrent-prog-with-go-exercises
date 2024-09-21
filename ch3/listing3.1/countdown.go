package main

import (
	"fmt"
	"time"
)

func main() {
	count := 5 // allocates memory space for an integer variable
	go countdown(&count) // shares memory by passing value of variable reference

	for count > 0 {
		time.Sleep(700 * time.Millisecond)
		fmt.Println(count)
	}
}

func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		*seconds -= 1
	}
}