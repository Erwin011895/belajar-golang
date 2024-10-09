package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/Erwin011895/belajar-golang/belajar_concurrency/topic"
)

// Global variable to package main
var URLs = []string{
	"https://jsonplaceholder.typicode.com/posts",    // 100
	"https://jsonplaceholder.typicode.com/comments", // 500
	"https://jsonplaceholder.typicode.com/albums",   // 100
}

func main() {
	topic.PlayDefer()
	fmt.Println("main goroutine start")
	defer fmt.Println("main goroutine end") // close program

	topic.BasicGoroutine()
	// topic.BasicChan()

	// fetchURL(URLs[0])

	// fetchSequentially()
	// fmt.Println("")
	// fetchConcurrently()
	// fmt.Println("")
	// fetchConcurrentlyWaitgroup()
}

// fetchURL launch HTTP call GET to url and return response
func fetchURL(url string) (string, error) {
	// Start recording the time when this function is executed
	start := time.Now()

	// Send an HTTP GET request to the given URL
	resp, err := http.Get(url)
	if err != nil {
		// If there is an error while sending the request, return the error
		return "", fmt.Errorf("error fetching %s: %v", url, err)
	}
	// Close resp.Body after the function is finished executing
	defer resp.Body.Close()

	// Read the entire response body from the server
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// If there is an error while reading the response body, return the error
		return "", fmt.Errorf("error reading response body from %s: %v", url, err)
	}

	// Calculate the duration it took to fetch the data from the URL
	duration := time.Since(start).Seconds()

	// Return the number of bytes fetched, the URL, and the duration it took
	fmt.Printf("Fetched %d bytes from %s in %.2f seconds\n", len(body), url, duration)
	return string(body), nil
}

func fetchSequentially() {
	fmt.Println("fetchSequentially")
	start := time.Now()
	// fmt.Println(start)

	var results []string

	// Iterate over each URL and fetch them one by one
	for _, url := range URLs {
		result, err := fetchURL(url)
		if err != nil {
			// Append the error message to results if an error occurs
			results = append(results, err.Error())
		} else {
			// Append the result to results if successful
			results = append(results, result)
		}
	}
	// fmt.Println("results", results[0]) // posts
	// fmt.Println("results", results[1]) // comments
	// fmt.Println("results", results[2]) // albums

	duration := time.Since(start).Seconds()
	fmt.Println("diff:", duration, "second")
}

func fetchConcurrently() {
	fmt.Println("fetchConcurrently")
	start := time.Now()
	// fmt.Println(start)

	// Create a channel to communicate results from goroutines
	ch := make(chan string)

	// Iterate over each URL
	for _, url := range URLs {
		// Launch a goroutine to fetch each URL concurrently
		go func(url string) {
			result, err := fetchURL(url)
			if err != nil {
				// Send the error message to the channel if an error occurs
				ch <- err.Error()
			} else {
				// Send the result to the channel if successful
				ch <- result
			}
		}(url)
	}

	var results []string
	// Collect results from the channel for each URL
	for range URLs { // iterate len(URLs) times without assign any new variable.
		results = append(results, <-ch)
	}

	// do something with results

	duration := time.Since(start).Seconds()
	fmt.Println("diff:", duration, "second")
}

func fetchConcurrentlyWaitgroup() {
	fmt.Println("fetchConcurrentlyWaitgroup")
	start := time.Now()
	// fmt.Println(start)

	// Create a channel to communicate results from goroutines
	// must be buffered channel.
	// ref: https://medium.com/goturkiye/concurrency-in-go-channels-and-waitgroups-25dd43064d1
	ch := make(chan string, 10)
	wg := sync.WaitGroup{}

	// Iterate over each URL
	for _, url := range URLs {
		// Launch a goroutine to fetch each URL concurrently
		wg.Add(1) // add 1 ke counter
		go func(url string) {
			defer wg.Done()
			result, err := fetchURL(url)
			if err != nil {
				// Send the error message to the channel if an error occurs
				ch <- err.Error()
			} else {
				// Send the result to the channel if successful
				ch <- result
			}
		}(url)
	}

	var results []string
	wg.Wait() // block until wg "counter" = zero

	// Collect results from the channel for each URL

	close(ch)             // close ch to ensure program can quit loop below
	for msg := range ch { // iterate len(URLs) times without assign any new variable.
		results = append(results, msg)
	}

	// do something with results

	duration := time.Since(start).Seconds()
	fmt.Println("diff:", duration, "second")
}
