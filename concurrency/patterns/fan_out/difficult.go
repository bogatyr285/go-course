package main

func main() {
apiUrls := []string{
	"https://jsonplaceholder.typicode.com/posts/1",
	"https://jsonplaceholder.typicode.com/posts/2",
	// Add more URLs as needed
}
numWorkers := 3 // Number of worker channels

// Create worker channels and a wait group for the workers
workerChans := make([]chan string, numWorkers)
for i := range workerChans {
	workerChans[i] = make(chan string)
}
var wg sync.WaitGroup
responseCh := make(chan APIResponse)

// Start worker goroutines
for i := 0; i < numWorkers; i++ {
	wg.Add(1)
	go worker(workerChans[i], responseCh, &wg)
}

// Distribute URLs to worker channels in a round-robin fashion
go func() {
	for i, url := range apiUrls {
		workerChans[i%numWorkers] <- url
	}
	for i := 0; i < numWorkers; i++ {
		close(workerChans[i]) // Close each worker channel when done
	}
}()

// Close the response channel once all workers are done
go func() {
	wg.Wait()
	close(responseCh)
}()

// Fan-in: Collect the results from the response channel
responses := aggregateResponses(responseCh)
for _, response := range responses {
	fmt.Printf("Title: %s, Body: %s\n", response.Title, response.Body)
}
// worker fetches responses from the given worker channel and sends them to the response channel
func worker(urlCh <-chan string, responseCh chan<- APIResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urlCh {
		// Simulate fetching from API
		resp := fetchFromAPI(url)
		responseCh <- resp
	}
}

// fetchFromAPI simulates an API call and returns an APIResponse
func fetchFromAPI(url string) APIResponse {
	return APIResponse{Title: fmt.Sprintf("Title from %s", url), Body: "Body of the response"}
}

// aggregateResponses collects responses from the response channel until the channel is closed
func aggregateResponses(ch <-chan APIResponse) []APIResponse {
	var responses []APIResponse
	for response := range ch {
		responses = append(responses, response)
	}
	return responses
}
}