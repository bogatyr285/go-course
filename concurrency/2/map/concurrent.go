package main

import (
	"fmt"
	"sync"
)

func main() {
	var myMap sync.Map

	// Function to add values to the map
	addFunc := func(key, value string, wg *sync.WaitGroup) {
		defer wg.Done()
		myMap.Store(key, value)
		fmt.Printf("Stored %s: %s\n", key, value)
	}

	// Function to read values from the map
	readFunc := func(key string, wg *sync.WaitGroup) {
		defer wg.Done()
		if value, ok := myMap.Load(key); ok {
			fmt.Printf("Loaded %s: %s\n", key, value)
		} else {
			fmt.Printf("Could not find key: %s\n", key)
		}
	}

	var wg sync.WaitGroup

	// Start multiple goroutines to store values
	wg.Add(3)
	go addFunc("key1", "value1", &wg)
	go addFunc("key2", "value2", &wg)
	go addFunc("key3", "value3", &wg)

	// Wait for storage operations to complete
	wg.Wait()

	// Start multiple goroutines to read values
	wg.Add(3)
	go readFunc("key1", &wg)
	go readFunc("key2", &wg)
	go readFunc("key3", &wg)

	// Wait for read operations to complete
	wg.Wait()

	// Update a value concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		myMap.Store("key2", "new_value2")
		fmt.Println("Updated key2 to new_value2")
	}()

	// Read the updated value
	wg.Add(1)
	go readFunc("key2", &wg)

	// Wait for all operations to complete
	wg.Wait()
}
