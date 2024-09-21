package main

import (
	"fmt"
	"time"
)

func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 1
	}
	fmt.Println("Stingy done")
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 1
	}
	fmt.Println("Spendy done")
}

func main() {
	money := 100
	go stingy(&money)
	go spendy(&money)
	time.Sleep(2 * time.Second) 
	fmt.Println("Money in bank account:", money)
}