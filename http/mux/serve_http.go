package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type PoliteServer struct {
}

func (ms *PoliteServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch {
	case req.Method == http.MethodPost && req.URL.Path == "/task":
		CreateTask(w, req)
		return
	case req.Method == http.MethodGet && req.URL.Path == "/task/{taskId}":
		GetTask(w, req)
		return
	}

	fmt.Fprintf(w, "Greetings! Thanks for visiting!\n")
}

func GetMemes(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Here you go!\n")
}

func main() {
	ps := &PoliteServer{}
	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	log.Println("before listen")

	go func() {
		err := http.ListenAndServe(":8090", ps)
		if err != nil {
			log.Fatalln(err)
		}
	}()
	<-ctx.Done()

	log.Println("after listen")
}

type LoggingMiddleware struct {
	handler http.Handler
}

func (lm *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	lm.handler.ServeHTTP(w, req)
	log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
}
