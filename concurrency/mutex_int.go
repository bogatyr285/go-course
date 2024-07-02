package main

import (
	"fmt"
	"sync"
)

func main() {

	value, x := 0, 5
	var wg sync.WaitGroup
	var mu sync.Mutex

	decrement := func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		value -= x
	}

	increment := func() {
		defer wg.Done()
		mu.Lock()
		value += x
		mu.Unlock()
	}

	for i := 0; i < 200; i++ {
		wg.Add(2)
		go increment()
		go decrement()
	}

	wg.Wait()
	fmt.Println(value)

}
