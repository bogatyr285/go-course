package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var buffer []int
	const bufferSize = 10

	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	var wg sync.WaitGroup

	produce := func(id int) {
		defer wg.Done()
		for i := 0; i < 10; i++ { // Limit the number of items produced for example purposes.
			mutex.Lock()
			for len(buffer) == bufferSize {
				cond.Wait() // Buffer is full
			}
			buffer = append(buffer, i)
			fmt.Printf("Producer %d produced %d\n", id, i)
			cond.Signal() // Signal consumers
			mutex.Unlock()
			time.Sleep(time.Millisecond * 500) // Simulate production time
		}
	}

	consume := func(id int) {
		defer wg.Done()
		for i := 0; i < 10; i++ { // Limit the number of items consumed for example purposes.
			mutex.Lock()
			for len(buffer) == 0 {
				cond.Wait() // Buffer is empty
			}
			item := buffer[0]
			buffer = buffer[1:]
			fmt.Printf("Consumer %d consumed %d\n", id, item)
			cond.Signal() // Signal producers
			mutex.Unlock()
			time.Sleep(time.Millisecond * 500) // Simulate consumption time
		}
	}

	const numProducers = 3
	const numConsumers = 3

	// Add to WaitGroup for producers and consumers
	wg.Add(numProducers + numConsumers)

	// Start producers
	for i := 1; i <= numProducers; i++ {
		go produce(i)
	}

	// Start consumers
	for i := 1; i <= numConsumers; i++ {
		go consume(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
