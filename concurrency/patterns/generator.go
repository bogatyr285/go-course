package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for val := range producer() {
		fmt.Println(val)
	}
}

func producer() chan int {
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
