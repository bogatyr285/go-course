package main

import (
	"sync"
	"testing"
)

func BenchmarkPubSub_Subscribe(b *testing.B) {
	ps := NewPubSub[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ps.Subscribe()
	}
}

// For a more realistic test, combining both subscribing and publishing
func BenchmarkPubSub_SubscribeAndPublish(b *testing.B) {
	ps := NewPubSub[int]()
	subscriberCount := 1000
	value := 1337

	var wg sync.WaitGroup

	// Create subscribers and start goroutines to consume messages
	for i := 0; i < subscriberCount; i++ {
		ch := ps.Subscribe()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range ch {
				// Consume the message; or you can process it
			}
		}()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ps.Publish(value)
	}

	// Cleanup
	ps.mu.Lock()
	ps.closed = true
	for _, ch := range ps.subscribers {
		close(ch)
	}
	ps.mu.Unlock()

	wg.Wait()
}

func BenchmarkPubSubV2_Subscribe(b *testing.B) {
	ps := NewPubSubV2[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ps.Subscribe()
	}
}

// For a more realistic test, combining both subscribing and publishing
func BenchmarkPubSubV2_SubscribeAndPublish(b *testing.B) {
	ps := NewPubSubV2[int]()
	subscriberCount := 1000
	value := 1337

	var wg sync.WaitGroup

	// Create subscribers and start goroutines to consume messages
	for i := 0; i < subscriberCount; i++ {
		ch := ps.Subscribe()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range ch {
				// Consume the message; or you can process it
			}
		}()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ps.Publish(value)
	}

	// Cleanup
	ps.Close()

	wg.Wait()
}
