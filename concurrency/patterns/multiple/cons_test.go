package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func runBenchmark(b *testing.B, noOfProducer int, noOfConsumer int) {
	for n := 0; n < b.N; n++ {
		data := make(chan int64)
		var producerGroup sync.WaitGroup
		var consumerCount int64

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
				for range data {
					atomic.AddInt64(&consumerCount, 1)
				}
			}(i)
		}
		consumerGroup.Wait()

		if atomic.LoadInt64(&ops) != atomic.LoadInt64(&consumerCount) {
			b.Fatalf("Mismatch: ops = %d, consumerCount = %d", ops, consumerCount)
		}
		// fmt.Println("processed ", atomic.LoadInt64(&ops))
	}
}

func BenchmarkSingleProducerConsumer(b *testing.B) {
	runBenchmark(b, 1, 1)
}

func BenchmarkMultipleProducersConsumers10(b *testing.B) {
	runBenchmark(b, 10, 10)
}

func BenchmarkMultipleProducersConsumers100(b *testing.B) {
	runBenchmark(b, 100, 100)
}
