package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// URL of the server
	url := "http://localhost:8081/"

	// Number of requests to perform
	numRequests := 30

	var wg sync.WaitGroup
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go func(requestNum int) {
			defer wg.Done()

			startTime := time.Now()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Request %d failed with error: %s\n", requestNum+1, err)
				return
			}
			defer resp.Body.Close()

			elapsedTime := time.Since(startTime)

			if resp.StatusCode == http.StatusOK {
				fmt.Printf("Request %d completed successfully in %s\n", requestNum+1, elapsedTime)
			} else {
				fmt.Printf("Request %d failed with response code %d\n", requestNum+1, resp.StatusCode)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All requests completed")
}
