package main

import (
	"fmt"
	"sync"
	"time"
)

type UserProfile struct {
	Name string
	Age  int
}

type UserStore struct {
	profiles map[string]UserProfile
	mu       sync.RWMutex
}

func NewUserStore() *UserStore {
	return &UserStore{
		profiles: make(map[string]UserProfile),
	}
}

func (store *UserStore) GetUser(username string) (UserProfile, bool) {
	store.mu.RLock()
	user, exists := store.profiles[username]
	store.mu.RUnlock()
	return user, exists
}

func (store *UserStore) AddUser(username string, profile UserProfile) {
	store.mu.Lock()
	store.profiles[username] = profile
	store.mu.Unlock()
}

func main() {
	store := NewUserStore()

	// Adding sample profiles
	store.AddUser("alice", UserProfile{Name: "Alice", Age: 30})
	store.AddUser("bob", UserProfile{Name: "Bob", Age: 25})

	var wg sync.WaitGroup

	// Simulate concurrent reads
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			if user, exists := store.GetUser("alice"); exists {
				fmt.Printf("Read Alice: %+v\n", user)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Simulate concurrent writes
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			store.AddUser("alice", UserProfile{Name: "Alice Updated", Age: 30 + i})
			fmt.Println("Updated Alice")
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()

	// Final read to check updates
	if user, exists := store.GetUser("alice"); exists {
		fmt.Printf("Final Alice: %+v\n", user)
	}
}
