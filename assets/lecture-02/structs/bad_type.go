package main

import (
	"fmt"
)

type S struct {
	a, b, c string
	m       map[string]string
}

func main() {
	x := S{"a", "b", "c", map[string]string{}}
	y := S{"a", "b", "c", map[string]string{}}

	fmt.Println(x == y)
}
