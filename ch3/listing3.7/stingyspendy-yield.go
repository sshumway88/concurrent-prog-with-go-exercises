package main

import (
	"fmt"
	"time"
)

func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 1
		// runtime.Gosched() // yields the processor, but does not suspend goroutine
	}
	fmt.Println("Stingy done")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 1
		// runtime.Gosched()
	}
	fmt.Println("Spendy done")
}

func main() {
	// runtime.GOMAXPROCS(1) // sets usesable CPUs to 1, and goroutines are user level threads, so they won't be pre-empted and execution will give expected result
	money := 100
	go stingy(&money)
	go spendy(&money)
	time.Sleep(3 * time.Second) 
	fmt.Println("Money in bank account:", money)
}