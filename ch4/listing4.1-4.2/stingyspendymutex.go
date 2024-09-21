package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money += 1
		mutex.Unlock()
	}
	fmt.Println("Stingy done")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money -= 1
		mutex.Unlock()
	}
	fmt.Println("Spendy done")
}

func main() {
	money := 100
	go stingy(&money)
	go spendy(&money)
	time.Sleep(2 * time.Second)
	// unlikely to need mutex after sleep
	mutex.Lock() 
	fmt.Println("Money in bank account:", money)
	mutex.Unlock()
}