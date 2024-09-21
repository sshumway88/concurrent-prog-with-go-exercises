package main

import (
	"fmt"
	"log"
	"os"

	"sync"
)

func readFile(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("failed to open file:", fileName)
	}

	fmt.Println(string(file))
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatal("must provide three file names")
		return
	}

	wg := &sync.WaitGroup{}
	for _, fName := range args {
		wg.Add(1)
		go readFile(fName, wg)
	}
	wg.Wait()
}
