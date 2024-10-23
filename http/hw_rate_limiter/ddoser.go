package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
	serverURL         = "http://localhost:7777" // Change this to your correct server address
	clientCount       = 5                       // Number of clients
	requestsPerClient = 20                      // Number of requests each client will send
	requestInterval   = 100 * time.Millisecond  // Interval between requests
)

func main() {
	var wg sync.WaitGroup

	// Launch multiple clients
	for i := 0; i < clientCount; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			for j := 0; j < requestsPerClient; j++ {
				makeRequest(clientID, j)
				time.Sleep(requestInterval)
			}
		}(i)
	}

	wg.Wait()
}

func makeRequest(clientID, requestID int) {
	resp, err := http.Get(serverURL)
	if err != nil {
		fmt.Printf("Client %d: Request %d failed: %v\n", clientID, requestID, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Client %d: Request %d failed to read response body: %v\n", clientID, requestID, err)
		return
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Printf("Client %d: Request %d rate limited: %s\n", clientID, requestID, string(body))
	} else {
		fmt.Printf("Client %d: Request %d succeeded: %s\n", clientID, requestID, string(body))
	}
}
