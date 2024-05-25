package main

import "fmt"

func main() {
	var a1 [10]byte

	var a [10]int
	fmt.Printf("Original array: %v\n", a)

	for i := 0; i < len(a1); i++ {
		a[i] = i * 2
	}

	fmt.Printf("Modified array: %v\n", a)
}
