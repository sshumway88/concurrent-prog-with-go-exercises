package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	go sayHello()
	// yields main goroutine to give go sayHello() chance to run, 
	// but no guarantee Go Scheduler won't pickup main goroutine again
	runtime.Gosched() 
	fmt.Println("Finished")
}
