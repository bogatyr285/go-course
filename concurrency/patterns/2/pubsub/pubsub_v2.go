package main

import (
	"fmt"
	"sync"
)

type PubSubV2[T any] struct {
	subscribers map[chan T]struct{}
	mu          sync.RWMutex
	closed      bool
	done        chan struct{}
}

func NewPubSubV2[T any]() *PubSubV2[T] {
	return &PubSubV2[T]{
		subscribers: make(map[chan T]struct{}),
		done:        make(chan struct{}),
	}
}

func (s *PubSubV2[T]) Subscribe() chan T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return nil
	}

	ch := make(chan T, 10) // Using buffered channel to avoid blocking
	s.subscribers[ch] = struct{}{}
	return ch
}

func (s *PubSubV2[T]) Unsubscribe(ch chan T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.subscribers[ch]; ok {
		delete(s.subscribers, ch)
		close(ch)
	}
}

func (s *PubSubV2[T]) Publish(value T) {
	s.mu.RLock()

	// If the PubSubV2 service is closed, avoid publishing
	if s.closed {
		s.mu.RUnlock()
		return
	}

	// Sending the value to all subscribers in a non-blocking way
	var wg sync.WaitGroup
	for ch := range s.subscribers {
		wg.Add(1)
		go func(ch chan T) {
			defer wg.Done()
			select {
			case ch <- value:
			case <-s.done:
				// Exit if the PubSubV2 is done
			}
		}(ch)
	}

	s.mu.RUnlock()
	wg.Wait()
}

func (s *PubSubV2[T]) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return
	}
	s.closed = true
	close(s.done)

	// Close all subscriber channels
	for ch := range s.subscribers {
		close(ch)
	}

	// Clear the subscribers map
	s.subscribers = nil
}

func main() {
	ps := NewPubSubV2[string]()

	wg := sync.WaitGroup{}

	for i := range 3 {
		sub := ps.Subscribe()
		go func() {
			wg.Add(1)

			for {
				select {
				case val, ok := <-sub:
					if !ok {
						fmt.Printf("sub %d, exiting\n", i)
						wg.Done()
						return
					}

					fmt.Printf("sub %d, value %v \n", i, val)
				}
			}
		}()
	}

	ps.Publish("one")
	ps.Publish("two")
	ps.Publish("three")

	ps.Close()

	wg.Wait()

	fmt.Println("completed")
}
