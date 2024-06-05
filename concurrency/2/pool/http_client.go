package main

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	url     = "http://localhost:18080/"
	numReqs = 10 // Adjust the number of requests as needed
)

func main() {
	var wg sync.WaitGroup
	wg.Add(numReqs) // Set the number of goroutines

	client := &http.Client{}

	for i := 0; i < numReqs; i++ {
		go func(i int) {
			defer wg.Done()

			resp, err := client.Get(url)
			if err != nil {
				fmt.Printf("Request %d failed: %v\n", i, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Request %d completed with status: %s\n", i, resp.Status)
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}
