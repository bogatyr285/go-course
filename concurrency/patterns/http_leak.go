package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result := make(chan struct {
		result any
		err    error
	})

	go func() {
		res, err := handler(ctx)
		result <- struct {
			result any
			err    error
		}{res, err}
	}()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err().Error())
	case res := <-result:
		fmt.Println(res)
	}

}

func handler(ctx context.Context) (string, error) {
	time.Sleep(time.Second * 2)
	log.Println("ready to return")
	return "handler", nil
}
