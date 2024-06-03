package main

import (
	"log"
	"time"
)

var x int64 = 0

func decrement() {
	for {
		x = x - 50

	}
}

func increment() {
	for {
		x = x + 15

	}
}

// run with -race flag
func main() {

	go increment()
	go decrement()
	time.Sleep(1 * time.Second)
	log.Println(x)
}
