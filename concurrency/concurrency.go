package main

import (
	"fmt"
	"time"
)

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(websiteChecker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	now := time.Now()

	fmt.Println("running for loop that runs goroutines")
	for index, url := range urls {
		fmt.Printf("calling anonymous goroutine #%v\n", index)
		go func(u string, index int) {
			fmt.Printf("putting stuff in channel for url %v\n", u)
			wcResult := websiteChecker(u)
			doneAt := time.Now()
			diff := doneAt.Sub(now)
			fmt.Printf("executing goroutine at index %v took %v\n", index, diff)
			resultChannel <- result{u, wcResult}
		}(url, index)
	}

	fmt.Println("running loop that reads from channel")
	for index := range urls {
		fmt.Printf("receiving from channel at #%v\n", index)
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
