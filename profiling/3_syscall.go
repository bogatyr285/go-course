package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"syscall"
	"time"
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
	for i := 0; i < 1000; i++ {
		fileName := fmt.Sprintf("/tmp/file_%d", time.Now().UnixNano())
		err := os.WriteFile(fileName, []byte("test"), 0644)
		if err != nil {
			log.Fatal(err)
		}
		err = syscall.Chmod(fileName, 0444)
		if err != nil {
			log.Fatal(err)
		}
	}
}
