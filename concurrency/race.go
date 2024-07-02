package main

import (
	"log"
	"sync/atomic"
	"time"
)

var x int64 = 0

func decrement() {
	// for {
	// x = x - 50
	atomic.AddInt64(&x, -50)
	// }
}

func increment() {
	for {
		atomic.AddInt64(&x, 15)
		// x = x + 15
	}
}

// run with -race flag
func main() {

	go increment()
	go decrement()
	time.Sleep(1 * time.Second)

	log.Println(atomic.LoadInt64(&x))
}
