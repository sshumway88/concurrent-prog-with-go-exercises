package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func readFile(searchStr string, fileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("failed to open file:", fileName)
	}

	switch  {
	case strings.Contains(string(file), searchStr):
		fmt.Println(fileName, "contains a match with", searchStr)	
	default:	
		fmt.Println(fileName, "does NOT contain a match with", searchStr)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 4 {
		log.Fatal("must provide three file names")
		return
	}

	searchStr := args[0]
	wg := &sync.WaitGroup{}
	for _, fName := range args[1:] {
		wg.Add(1)
		go readFile(searchStr, fName, wg)
	}
	wg.Wait()
}
