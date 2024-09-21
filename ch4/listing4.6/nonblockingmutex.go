package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, lock *sync.Mutex) {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic("http call error")
	}

	body, _ := io.ReadAll(resp.Body)
	// entire for loop runs very quickly, so better to manipulate lock outside of it rather than around each write
	lock.Lock()
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			// There is a cost to frequently calling Lock/Unlock. Will slow down performance, so want to weigh
			// frequency of lock manipulation against speed of code section being protected
			// lock.Lock(). Minimize the amount of time spent holding mutex locks, while also trying to lower the number of mutex calls
			frequency[cIndex] += 1
			// lock.Unlock()
		}
	}
	lock.Unlock()
	fmt.Println("Completed:", url, time.Now().Format(time.TimeOnly))
}

func main() {
    mutex := sync.Mutex{}
    var frequency = make([]int, 26)
    for i := 2000; i <= 2200; i++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        go countLetters(url, frequency, &mutex)
    }
    for i := 0; i < 100; i++ {
        time.Sleep(100 * time.Millisecond)
        if mutex.TryLock() {
            for i, c := range allLetters {
                fmt.Printf("%c-%d ", c, frequency[i])
            }
            fmt.Println()
            mutex.Unlock()
        } else {
            fmt.Println("Mutex already being used")
        }
    }
}