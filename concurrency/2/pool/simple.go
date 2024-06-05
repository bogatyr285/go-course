package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			time.Sleep(2 * time.Second)
			fmt.Println("creating new object")
			return struct{}{}
		},
	}

	// this will invoke New
	obj := pool.Get()
	pool.Put(obj)

	// this will not invoke the New func
	pool.Get()
}
