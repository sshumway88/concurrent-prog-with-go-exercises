package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

func countdown(seconds *int) {
	countdown := *seconds
    for countdown > 0 {
        time.Sleep(1 * time.Second)
		lock.Lock()
        *seconds -= 1
		countdown = *seconds
		lock.Unlock()
    }
}

func main() {
    count := 5
	localCount := count
    go countdown(&count)
    for localCount > 0 {
        time.Sleep(500 * time.Millisecond)
		lock.Lock()
        fmt.Println(count)
		localCount = count
		lock.Unlock()
    }
}