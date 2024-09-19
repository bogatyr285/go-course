package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func leakyFunction() {
	ch := make(chan int)
	// var data [][]byte
	go func() {
		for {
			select {
			case <-ch:
				return
			default:
				// Simulate work
				_ = make([]byte, 1024*1024)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
}

// curl http://localhost:6060/debug/pprof/heap\?seconds\=30 > heap.prof
// go tool pprof -http=":" heap.prof
// (pprof) top
//
//	(pprof) list leakyFunction
//	(pprof) web
func main() {
	for i := 0; i < 100; i++ {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)

		leakyFunction2(ctx)
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		runtime.GC()
		var m runtime.MemStats

		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %d", m.Alloc/1024/1024)
		fmt.Printf("\tTotalAlloc = %v MB", m.TotalAlloc/1024/1024)
		fmt.Printf("\tSys = %v MB", m.Sys/1024/1024)
		fmt.Printf("\tNumGC = %v\n", m.NumGC)
		time.Sleep(5 * time.Second)
	}
}

//

func leakyFunction2(ctx context.Context) {

	go func() {
		defer fmt.Println("done")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_ = make([]byte, 1024*1024)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}
