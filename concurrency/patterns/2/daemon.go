package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
)

// PeriodicTask represents a struct which handles the periodic execution of a function
type PeriodicTask struct {
	period time.Duration
	task   func()
}

func New(period time.Duration, task func()) *PeriodicTask {
	return &PeriodicTask{
		period: period,
		task:   task,
	}
}

func (pt *PeriodicTask) Run(ctx context.Context) {
	ticker := time.NewTicker(pt.period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			pt.task()
		case <-ctx.Done():
			log.Println("Stopping periodic task:", ctx.Err())
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	task := func() {
		log.Println("Executing periodic task at:", time.Now())
	}

	pt := New(time.Second*2, task)
	go pt.Run(ctx)

	<-ctx.Done()
	log.Println("Daemon has been stopped")
}
