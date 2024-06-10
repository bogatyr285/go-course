package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Number of workers in the pool
const workerCount = 5

// outputDir is the directory to save downloaded files
const outputDir = "./downloads"

// DownloadTask represents a file download task
type DownloadTask struct {
	URL      string
	FileName string
}

// downloadFile performs the file download
func downloadFile(task DownloadTask) error {
	resp, err := http.Get(task.URL)
	if err != nil {
		return fmt.Errorf("failed to download %s: %v", task.URL, err)
	}
	defer resp.Body.Close()

	// Create output file
	out, err := os.Create(fmt.Sprintf("%s/%s", outputDir, task.FileName))
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", task.FileName, err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write%v to file %s: %v", task.URL, task.FileName, err)
	}
	return nil
}

// worker function processes download tasks from the input channel and sends results to the output channel
func worker(wg *sync.WaitGroup, tasks <-chan DownloadTask, results chan<- error) {
	defer wg.Done()
	for task := range tasks {
		err := downloadFile(task)
		results <- err
	}
}

func main() {
	// Ensure output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create output directory: %v\n", err)
			return
		}
	}

	// List of download tasks
	tasks := []DownloadTask{
		{URL: "https://example.com/file1.jpg", FileName: "file1.jpg"},
		{URL: "https://example.com/file2.jpg", FileName: "file2.jpg"},
		{URL: "https://example.com/file3.jpg", FileName: "file3.jpg"},
		{URL: "https://example.com/file4.jpg", FileName: "file4.jpg"},
		{URL: "https://example.com/file5.jpg", FileName: "file5.jpg"},
	}

	taskChannel := make(chan DownloadTask, len(tasks))
	resultChannel := make(chan error, len(tasks))

	// Start the timer
	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, taskChannel, resultChannel)
	}

	// Send tasks to workers
	for _, task := range tasks {
		taskChannel <- task
	}
	close(taskChannel)

	// Wait for all workers to finish
	wg.Wait()
	close(resultChannel)

	// Check results
	for err := range resultChannel {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	// End the timer
	elapsedTime := time.Since(startTime)
	fmt.Printf("All downloads completed in %s\n", elapsedTime)
}
