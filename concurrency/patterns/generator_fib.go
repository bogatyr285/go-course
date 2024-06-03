package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for c := range fibonacci(1000000) {
		fmt.Printf("Fibonacci number is %d \n", c)
	}
}

func fibonacci(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		fmt.Println("Hello, I am the producer")
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}
