package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic("http call error")
	}

	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed:", url)
}

func main() {
	frequency := make([]int, 26)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency)			
	}

	time.Sleep(10 * time.Second)
	for i, c := range allLetters {
		fmt.Printf("%c-%d\n", c, frequency[i])
	}
}