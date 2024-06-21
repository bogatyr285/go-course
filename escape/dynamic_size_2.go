package main

func main() {
	foo(10)
}

func foo(cap int) {
	a := make([]string, 0, cap)
	println(a)
}
