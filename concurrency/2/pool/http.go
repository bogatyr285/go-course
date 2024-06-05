package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var bufferPool = &sync.Pool{
	New: func() interface{} {
		log.Println("init buffer")
		return new(bytes.Buffer)
	},
}

// Handler function that uses the buffer pool
func requestHandler(w http.ResponseWriter, r *http.Request) {
	// Get a buffer from the pool
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)

	// Reset the buffer before use
	buf.Reset()

	// Read the request body into the buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Process the buffer content (for example, just print it here)
	fmt.Println("Received request:", buf.String())

	// Write a response
	w.Write([]byte("Request received successfully"))
}

func main() {
	http.HandleFunc("/", requestHandler)
	fmt.Println("Server started at :8080")
	log.Fatalln(http.ListenAndServe(":18080", nil))
}
