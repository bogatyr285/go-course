package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runTime = 10 * time.Second

	greedyRoutine1 := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runTime; {
			sharedLock.Lock()
			time.Sleep(10 * time.Nanosecond)
			sharedLock.Unlock()
			runtime.Gosched()
			count++
		}
		fmt.Printf("Greedy worker 1 was able to execute %v work loops\n", count)
	}

	greedyRoutine2 := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runTime; {
			sharedLock.Lock()
			time.Sleep(5 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(5 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		fmt.Printf("Greedy worker 2 was able to execute %v work loops\n", count)
	}

	nonGreedyRoutine := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runTime; {
			for i := 0; i < 10; i++ {
				sharedLock.Lock()
				time.Sleep(1 * time.Nanosecond)
				sharedLock.Unlock()
			}
			count++
		}
		fmt.Printf("Non greedy worker was able to execute %v work loops.\n", count)
	}

	wg.Add(3)
	go greedyRoutine1()
	go greedyRoutine2()
	go nonGreedyRoutine()
	wg.Wait()
}
