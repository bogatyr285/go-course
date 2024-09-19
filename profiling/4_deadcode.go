package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("profile.pb.gz")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	for range 100000 {
		fmt.Println(multiply(1, 9))
	}
}

func multiply(x, y int) int {
	a := x + 1
	b := a + 2*y
	b = b * b
	return x * y
}
