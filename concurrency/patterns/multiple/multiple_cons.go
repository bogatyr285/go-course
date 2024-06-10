package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	noOfProducer := 10
	noOfConsumer := 10
	data := make(chan int64)
	var producerGroup sync.WaitGroup

	// producer
	var ops int64
	for i := 0; i < noOfProducer; i++ {
		producerGroup.Add(1)
		go func() {
			defer producerGroup.Done()
			for c := 0; c < 100; c++ {
				data <- atomic.AddInt64(&ops, 1)
			}
		}()
	}

	go func() {
		defer close(data)
		producerGroup.Wait()
	}()

	var consumerGroup sync.WaitGroup
	// consumer
	for i := 0; i < noOfConsumer; i++ {
		consumerGroup.Add(1)
		go func(i int) {
			defer consumerGroup.Done()
			for data := range data {
				fmt.Printf("Value of i = %d Printed by consumer %d\n", data, i)
			}
		}(i)
	}
	consumerGroup.Wait()
}
