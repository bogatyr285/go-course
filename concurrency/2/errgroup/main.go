package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func fetchURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 response from %s: %d", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body from %s: %v", url, err)
	}

	fmt.Printf("Fetched %d bytes from %s\n", len(body), url)
	return nil
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
	}

	var g errgroup.Group

	for _, url := range urls {
		url := url
		g.Go(func() error {
			return fetchURL(url)
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatalf("Some HTTP requests failed: %v\n", err)
		return
	}

	log.Println("All HTTP requests succeeded.")
}
