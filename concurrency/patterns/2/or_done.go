package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// попробуйте закрыть канал сразу же
	// тогда значения не выводятся
	// cancel()
	go func() {
		time.Sleep(300 * time.Millisecond)
		cancel()
	}()
	for val := range orDone(ctx, gen(1, 2, 3)) {
		fmt.Println("val:", val)
	}
}

func orDone[T any](ctx context.Context, in <-chan T) <-chan T {
	valStream := make(chan T)
	go func() {
		defer close(valStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-ctx.Done():
				}
			}
		}
	}()
	return valStream
}

func gen(nums ...int) chan int {
	out := make(chan int, len(nums))
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
			// закомментируй для теста
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return out
}
