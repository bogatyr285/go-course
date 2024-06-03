package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		time.Sleep(300*time.Millisecond)
		fmt.Println("go routine: done")
	}()
	wg.Wait()
	fmt.Println("executed immediately")
}
