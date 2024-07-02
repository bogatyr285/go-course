package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data int
	mu   sync.Mutex
	cond = sync.NewCond(&mu)
)

func producer() {
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(2 * time.Second) // Simulate work
	data = 42
	fmt.Println("Producer: Data produced")
	cond.Broadcast()

	// cond.Signal()
	// cond.Signal()
}

func consumer() {
	mu.Lock()
	defer mu.Unlock()
	cond.Wait() // Wait for the producer to signal
	fmt.Printf("Consumer: Data consumed: %d\n", data)
}

func consumer2() {
	mu.Lock()
	defer mu.Unlock()
	cond.Wait() // Wait for the producer to signal
	fmt.Printf("Consumer222: Data consumed: %d\n", data)
}

func main() {
	go consumer2()
	go consumer()

	go producer()

	time.Sleep(3 * time.Second) // Make sure to give enough time for goroutines to finish
}
