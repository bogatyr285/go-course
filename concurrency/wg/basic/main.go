package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("3")
	}()

	wg.Wait()
}
