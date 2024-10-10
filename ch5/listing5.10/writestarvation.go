package main

import (
	"fmt"
	"time"

	listing4_12 "github.com/sshumway/concurrent-prog-with-go/ch4/exercises/ex4.2_3"
)

func main() {
	rwMutex := listing4_12.ReadWriteMutex{}
	i := 0
	for range 2 {
		fmt.Println("i", i)
		i++
		go func() {
			for {
				rwMutex.ReadLock()
				time.Sleep(1 * time.Second)
				fmt.Println("read done")
				rwMutex.ReaderUnlock()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	rwMutex.WriteLock()
	fmt.Println("Write finished")
}
