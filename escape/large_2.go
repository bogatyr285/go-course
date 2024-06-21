package main

import "unsafe"

type article struct {
	name  string
	views int
}

func main() {
	v := new([2731]article)
	println("a's size:", unsafe.Sizeof(article{}))
	println("v's size:", unsafe.Sizeof(*v))
}
