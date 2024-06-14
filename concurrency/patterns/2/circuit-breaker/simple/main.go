package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cb := newCircuitBreaker(3, 4*time.Second)
	errs := []error{
		errors.New("err 1"),
		// errors.New("err 2"),
		// errors.New("err 3"),
		nil,
		nil,
		nil,
		nil,
	}
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		_, _ = cb.Execute(func() (interface{}, error) {
			// suppose this is a wrapped http call
			r := rand.Intn(len(errs))
			return []byte{}, errs[r]
		})
	}
}

type circuitBreaker struct {
	failures, maxFailures int
	open                  chan struct{}
	state                 string
}

func newCircuitBreaker(maxFailures int, timeout time.Duration) *circuitBreaker {
	cb := &circuitBreaker{
		maxFailures: maxFailures,
		open:        make(chan struct{}),
		state:       "closed",
	}
	go func() {
		for {
			select {
			case <-cb.open:
				fmt.Println("half-open. wait for ", timeout)
				select {
				case <-time.After(timeout):
					cb.state = "half-open"
					cb.failures = 0
				}
			}
		}
	}()
	return cb
}

func (c *circuitBreaker) Execute(fn func() (interface{}, error)) ([]byte, error) {
	openCircuit := func() {
		c.state = "open"
		select {
		case c.open <- struct{}{}:
		default:
		}
	}
	if c.failures >= c.maxFailures {
		fmt.Println("\t circuit: open: ", c.failures)
		c.state = "open"
	}
	if c.state == "half-open" && c.failures > 0 {
		fmt.Println("\t circuit: half-open ", c.failures)
		c.state = "open"
		openCircuit()
	}
	if c.state == "open" {
		fmt.Println("circuit: open")
		openCircuit()
		return []byte{}, nil
	}

	fmt.Println("circuit: closed")
	res, err := fn()
	if err != nil {
		fmt.Println("\t call failed: ", err)
		c.failures++
		return []byte{}, err
	}
	fmt.Println("\t request ok")
	c.state = "closed"
	return res.([]byte), nil
}
