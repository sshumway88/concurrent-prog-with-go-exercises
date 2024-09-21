package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func matchRecorder(matchEvents *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		*matchEvents = append(*matchEvents,
			"Match Event " + strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		// fmt.Println("Appended match event")
	}
}

func clientHandler(mEvents *[]string, mutex *sync.RWMutex, st time.Time) {
	for i := 0; i < 100; i++ {
		mutex.RLock()
		allEvents := copyAllEvents(mEvents)
		mutex.RUnlock()
		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	allEvents := make([]string, len(*matchEvents))
	allEvents = append(allEvents, *matchEvents...)
	return allEvents
}

func main() {
	mutex := &sync.RWMutex{}
	// create fake match events
	matchEvents := make([]string, 0, 5000)
	for j := 0; j < 10000; j++ {
		matchEvents = append(matchEvents, "match event")
	}

	go matchRecorder(&matchEvents, mutex)
	st := time.Now()
	for i := 0; i < 5000; i++ {
		go clientHandler(&matchEvents, mutex, st)
	}
	time.Sleep(10 * time.Second)
}
