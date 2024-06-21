package main

import "unsafe"

// https://go.dev/src/cmd/compile/internal/ir/cfg.go

type article struct {
	name  string
	views int
}

func main() {
	v := new([2730]article)
	println("a's size:", unsafe.Sizeof(""))
	println("v's size:", unsafe.Sizeof(*v))
}
