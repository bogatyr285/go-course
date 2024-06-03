package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("I am goroutine 1")
	}()

	go func() {
		fmt.Println("I am goroutine 2")
	}()
}
