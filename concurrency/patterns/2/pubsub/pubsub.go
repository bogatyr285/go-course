package main

import (
	"sync"
)

type PubSub[T any] struct {
	subscribers []chan T
	mu          sync.RWMutex
	closed      bool
}

func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{
		mu:          sync.RWMutex{},
		subscribers: make([]chan T, 0),
	}
}

func (s *PubSub[T]) Subscribe() <-chan T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return nil
	}

	r := make(chan T)

	s.subscribers = append(s.subscribers, r)

	return r
}

func (s *PubSub[T]) Publish(value T) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.closed {
		return
	}

	for _, ch := range s.subscribers {
		ch <- value
	}
}

func (s *PubSub[T]) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return
	}

	for _, ch := range s.subscribers {
		close(ch)
	}

	s.closed = true
}

// func main() {
// 	ps := NewPubSub[string]()

// 	wg := sync.WaitGroup{}

// 	for i := range 3 {
// 		sub := ps.Subscribe()
// 		go func() {
// 			wg.Add(1)

// 			for {
// 				select {
// 				case val, ok := <-sub:
// 					if !ok {
// 						fmt.Printf("sub %d, exiting\n", i)
// 						wg.Done()
// 						return
// 					}

// 					fmt.Printf("sub %d, value %v \n", i, val)
// 				}
// 			}
// 		}()
// 	}

// 	ps.Publish("one")
// 	ps.Publish("two")
// 	ps.Publish("three")

// 	ps.Close()

// 	wg.Wait()

// 	fmt.Println("completed")
// }
