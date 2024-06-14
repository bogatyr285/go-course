package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	listenAddr      = "127.0.0.1:8080"
	shutdownTimeout = 5 * time.Second
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runServer(ctx); err != nil {
		log.Fatal(err)
	}
}

func runServer(ctx context.Context) error {
	var (
		mux = http.NewServeMux()
		srv = &http.Server{
			Addr:    listenAddr,
			Handler: mux,
		}
		c = &Closer{}
	)

	mux.Handle("/", handleIndex())

	c.Add(srv.Shutdown)

	c.Add(func(ctx context.Context) error {
		time.Sleep(3 * time.Second)

		return nil
	})

	c.Add(func(ctx context.Context) error {
		time.Sleep(3 * time.Second)
		return errors.New("oops error occurred")
	})

	c.Add(func(ctx context.Context) error {
		return errors.New("uh-oh, another error occurred")
	})

	c.Add(func(ctx context.Context) error {
		return errors.New("uh-oh, another error occurred")
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %v", err)
		}
	}()

	log.Printf("listening on %s", listenAddr)
	<-ctx.Done()

	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := c.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer: %v", err)
	}

	return nil
}

func handleIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
}

// Closer

type Func func(ctx context.Context) error

type Closer struct {
	mu    sync.Mutex
	funcs []Func
}

func (c *Closer) Add(f Func) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, f)
}

func (c *Closer) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var (
		msgs     = make([]string, 0, len(c.funcs))
		complete = make(chan struct{}, 1)
	)

	go func() {
		for _, f := range c.funcs {
			if err := f(ctx); err != nil {
				msgs = append(msgs, fmt.Sprintf("[!] %v", err))
			}
		}

		complete <- struct{}{}
	}()

	select {
	case <-complete:
		break
	case <-ctx.Done():
		return fmt.Errorf("shutdown cancelled: %v", ctx.Err())
	}

	if len(msgs) > 0 {
		return fmt.Errorf(
			"shutdown finished with error(s): \n%s",
			strings.Join(msgs, "\n"),
		)
	}

	return nil
}

// func (c *Closer) CloseV2(ctx context.Context) error {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	var (
// 		wg      sync.WaitGroup
// 		mu      sync.Mutex
// 		msgs    = make([]string, 0)
// 		errChan = make(chan error, len(c.funcs))
// 	)

// 	// Launch each closer function in its own goroutine
// 	wg.Add(len(c.funcs))
// 	for _, f := range c.funcs {
// 		go func(f func(context.Context) error) {
// 			defer wg.Done()

// 			if err := f(ctx); err != nil {
// 				mu.Lock()
// 				msgs = append(msgs, fmt.Sprintf("[!] %v", err))
// 				mu.Unlock()
// 			}
// 		}(f)
// 	}

// 	// Wait for all goroutines to finish
// 	go func() {
// 		wg.Wait()
// 		close(errChan)
// 	}()

// 	// Monitor for context cancellation
// 	select {
// 	case <-ctx.Done():
// 		return fmt.Errorf("shutdown cancelled: %v", ctx.Err())
// 	case <-errChan:
// 		// Wait for the WaitGroup to be done
// 		wg.Wait()
// 	}

// 	if len(msgs) > 0 {
// 		return fmt.Errorf("shutdown finished with error(s): \n%s", strings.Join(msgs, "\n"))
// 	}

// 	return nil
// }
