package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"time"
	"sync"
)

func main() {
	store := make(map[string]int)
	mutex := &sync.RWMutex{}

	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, store, mutex)
	}

	time.Sleep(10 * time.Second)

	mutex.Lock()
	for k, v := range store {
		fmt.Printf("%s -> %d\n", k, v)
	}
	mutex.Unlock()
}

func countWords(url string, store map[string]int, mutex *sync.RWMutex) {
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

	// can put the mutex in loop, but apparently
	// locking and unlocking has their own overhead
	mutex.Lock()
	for scanner.Scan() {
		word := scanner.Text()

		if _, ok := store[word]; !ok {
			store[word] = 1
		} else {
			store[word]++
		}
	}
	mutex.Unlock()

	if err := scanner.Err(); err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
}