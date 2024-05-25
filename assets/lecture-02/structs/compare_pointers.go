package main

import (
	"fmt"
)

type S struct {
	a, b, c string
}

func main() {
	x := &S{"a", "b", "c"}
	y := &S{"a", "b", "c"}

	fmt.Println(&x, &y)
}
