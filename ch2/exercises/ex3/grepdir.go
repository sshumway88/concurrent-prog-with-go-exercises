package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func readFile(searchStr string, filePath string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("failed to open file:", filePath)
		return
	}

	switch {
	case strings.Contains(string(file), searchStr):
		fmt.Println(filePath, "contains a match with", searchStr)
	default:
		fmt.Println(filePath, "does NOT contain a match with", searchStr)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("must provide three file names")
		return
	}

	searchStr := args[0]
	dirPath := args[1]

	dirContents, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal("Failed to open file directory")
	}

	wg := &sync.WaitGroup{}
	for _, entry := range dirContents {
		wg.Add(1)
		go readFile(searchStr, filepath.Join(dirPath, entry.Name()), wg)
	}
	wg.Wait()
}
