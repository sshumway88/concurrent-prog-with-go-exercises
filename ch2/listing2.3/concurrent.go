package main

import (
	"fmt"
	"time"
)

func doWork(id int) {
	fmt.Printf("Work %d started at %s\n", id, time.Now().Format("15:04:05"))
    time.Sleep(1 * time.Second)
    fmt.Printf("Work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}

func main() {
	for i := 0; i < 5; i++ {
		// asynchronous call (don't have to wait for it to finish)
		// execution order of concurrent jobs can never be guaranteed
		// OS might pick executions in a different order than they were created
		go doWork(i)
	}
	time.Sleep(2 * time.Second)
}