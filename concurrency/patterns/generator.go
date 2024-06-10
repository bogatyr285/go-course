package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for val := range producer(wg) {
		fmt.Println(val)
	}
}

func producer(wg *sync.WaitGroup) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		defer wg.Wait()
		for i := 0; i < 100; i += 5 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				out <- i
			}(i)
		}
	}()
	return out
}
