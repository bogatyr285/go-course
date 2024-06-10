package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type APIResponse struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func fetchFromAPI(wg *sync.WaitGroup, url string, ch chan<- APIResponse) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request to %s: %v\n", url, err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response from %s: %v\n", url, err)
		return
	}

	var apiResponse APIResponse
	json.Unmarshal(body, &apiResponse)
	ch <- apiResponse
}

func aggregateResponses(ch <-chan APIResponse) []APIResponse {
	var responses []APIResponse
	for c := range ch {
		responses = append(responses, c)
	}
	return responses
}

func main() {
	var wg sync.WaitGroup
	apiUrls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/6",
	}
	ch := make(chan APIResponse, len(apiUrls))

	// Fan-out: Start a goroutine for each API request
	for _, url := range apiUrls {
		wg.Add(1)
		go fetchFromAPI(&wg, url, ch)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(ch)

	// Fan-in: Aggregate the responses
	responses := aggregateResponses(ch)

	// Print aggregated responses
	for _, response := range responses {
		fmt.Printf("Title: %s, Body: %s\n", response.Title, response.Body)
	}
}
