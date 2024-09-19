package main

import (
	"log"
	"os"
	"runtime/debug"
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
	firstFunctionToBeCalled()
	defer pprof.StopCPUProfile()
}

func firstFunctionToBeCalled() {
	secondFunctionToBeCalled()
}

func secondFunctionToBeCalled() {
	doSum()
	debug.PrintStack()
}

func doSum() int {
	sum := 0
	for i := 0; i < 787766777; i++ {
		sum += i
	}
	return sum
}
