package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// RateLimitedWriter limits the rate of writing
type RateLimitedWriter struct {
	writer          io.Writer
	rateBytesPerSec int
}

// Write writes data at a limited rate
func (rlw *RateLimitedWriter) Write(p []byte) (int, error) {
	start := time.Now()
	totalWritten := 0
	for totalWritten < len(p) {
		toWrite := len(p) - totalWritten
		if toWrite > rlw.rateBytesPerSec {
			toWrite = rlw.rateBytesPerSec
		}
		n, err := rlw.writer.Write(p[totalWritten : totalWritten+toWrite])
		if err != nil {
			return totalWritten, err
		}
		totalWritten += n
		sleepTime := time.Duration(float64(n) / float64(rlw.rateBytesPerSec) * float64(time.Second))
		time.Sleep(sleepTime)
	}
	elapsedTime := time.Since(start)
	fmt.Printf("Total bytes written: %d, took %v\n", totalWritten, elapsedTime)
	return totalWritten, nil
}
func main() {
	rateLimitedWriter := &RateLimitedWriter{
		writer:          os.Stdout,
		rateBytesPerSec: 10, // Limit the write rate to 10 bytes per second
	}
	data := "This is a test data to demonstrate rate limiting in Go.\n"
	rateLimitedWriter.Write([]byte(data))
}
