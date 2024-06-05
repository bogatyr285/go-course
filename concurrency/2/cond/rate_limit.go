package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	mu          = &sync.Mutex{}
	cond        = sync.NewCond(mu)
	activeTasks = 0
	maxTasks    = 5 // maximum number of concurrent tasks
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	for activeTasks >= maxTasks {
		cond.Wait() // Wait until the condition changes (i.e., some tasks complete)
	}
	activeTasks++
	mu.Unlock()

	defer func() {
		mu.Lock()
		activeTasks--
		cond.Signal() // Signal one waiting goroutine
		mu.Unlock()
	}()

	// Simulate a resource-intensive task
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Task Completed\n")
}

func main() {
	port := ":8081"
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
