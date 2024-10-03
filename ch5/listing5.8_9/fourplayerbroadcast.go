package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	playersInGame := 4
	for playerID := 0; playerID < 4; playerID++ {
		go playerHandler(cond, playerID, &playersInGame)
		time.Sleep(1 * time.Second)
	}
}

func playerHandler(cond *sync.Cond, playerID int, playersInGame *int) {
	cond.L.Lock()
	*playersInGame--
	if *playersInGame == 0 {
		cond.Broadcast()
	}
	if *playersInGame != 0 {
		fmt.Println(playerID, "is waiting for more player")
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println("All players are connected. Ready player", playerID)
}