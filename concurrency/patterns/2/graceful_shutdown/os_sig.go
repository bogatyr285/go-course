package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	exit := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	log.Println("app run")
	// what if we'll have huge init part?
	// dbs, metrics, grpc/http, kafka, etc?
	q := <-exit
	log.Println("app shut down: ", q)
}
