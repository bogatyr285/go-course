package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	wg.Done()
	wg.Wait()
}
