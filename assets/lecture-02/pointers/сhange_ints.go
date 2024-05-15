package main

import "fmt"

func main() {
	var y = 123
	var p = &y
	*p = 1337

	fmt.Println(y)
}
