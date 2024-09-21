package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// concurrent writes to this map causes a fatal error
var wordCounts = make(map[string]int)

func countWords(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Panic("http call error")
	}

	body, _ := io.ReadAll(resp.Body)
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	words := wordRegex.FindAllString(string(body), -1)
	
	for _, w := range words {
		lowered := strings.ToLower(w)
		wordCounts[lowered] += 1
	}
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make(map[string]int)
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url)
	}
	time.Sleep(10 * time.Second)
	for k, v := range frequency {
		fmt.Println(k, "->", v)
	}
}