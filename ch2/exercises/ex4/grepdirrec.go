package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func searchDir(searchStr, path string, dirInfo fs.DirEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	fullPath := filepath.Join(path, dirInfo.Name())
	if dirInfo.IsDir() {
		dirContents, err := os.ReadDir(fullPath)
		if err != nil {
			fmt.Println("failed to open directory:", dirInfo.Name())
		}

		for _, dirEntry := range dirContents {
			wg.Add(1)
			go searchDir(searchStr, fullPath, dirEntry, wg)
		}
		return
	} 
	readFile(searchStr, fullPath)
}

func readFile(searchStr string, filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("failed to open file:", filePath)
		return
	}

	switch  {
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
	for _, dirInfo := range dirContents {
		wg.Add(1)
		go searchDir(searchStr, dirPath, dirInfo, wg)
	}
	wg.Wait()
}
