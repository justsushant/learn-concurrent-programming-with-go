package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"time"
)

// received errors when converted to concurrent from sequential prog
// when run with, go run -race wordfrequency.go, it shows race conditions

func main() {
	store := make(map[string]int)

	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, store)
	}

	time.Sleep(10 * time.Second)

	for k, v := range store {
		fmt.Printf("%s -> %d\n", k, v)
	}
}

func countWords(url string, store map[string]int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Received error while getting page %s\n", url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Received Status Code %s for the url, %s\n", resp.Status, url)
		return
	}

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()

		if _, ok := store[word]; !ok {
			store[word] = 1
		} else {
			store[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
}