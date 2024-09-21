package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	
	// Go defaults value of GOMAXPROCS to number of CPU
	// passing n < 1 to GOMAXPROCS(n) returns current value without altering it
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}