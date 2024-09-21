package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func matchRecorder(matchEvents *[]string, mutex *sync.Mutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		*matchEvents = append(*matchEvents, 
			"Match event " + strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Appended match event")	
	}
}

func clientHandler(mEvents *[]string, mutex *sync.Mutex, st time.Time) {
	for i := 0; i < 100; i++ {
		mutex.Lock()
		allEvents := copyAllEvents(mEvents)
		mutex.Unlock()

		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(mEvents *[]string) []string {
	allEvents := make([]string, len(*mEvents))
	for _, e := range *mEvents {
		allEvents = append(allEvents, e)
	}
	return allEvents
}

func main() {
	mutex := &sync.Mutex{}
	matchEvents := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		matchEvents = append(matchEvents, "match event")
	}
	
	go matchRecorder(&matchEvents, mutex)
	
	start := time.Now()
	for i := 0; i < 5000; i++ {
		go clientHandler(&matchEvents, mutex, start)
	}
	time.Sleep(10 * time.Second)
}