package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	odd := OddIntGen(2)
	even := EvenIntGen(2)
	hex := HexIntGen(2)
	out := FanIn(done, odd, even, hex)

	for n := range out {
		fmt.Println("fanned number:", n)
	}
}

// FanIn считывает данные из нескольких каналов и записывает их в 1 конечный канал
// FAN-IN - функция принимает несколько каналов в качестве входных данных
// Она считывает все входные данные и отправляет все значения в 1 конечный выходной канал
func FanIn(done chan struct{}, inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(inputs)) // !

	for _, in := range inputs {
		go func(numbers <-chan int) {
			defer wg.Done()
			for n := range numbers {
				select {
				case <-done:
					return
				case out <- n:
				}
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// EvenIntGen generates random even integers
func EvenIntGen(total int) <-chan int {
	out := make(chan int)
	go func() {
		n := randInt(10000)
		for i := n; i > n-total*2; i-- {
			if i%2 == 0 {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

// OddIntGen generates random odd integers
func OddIntGen(total int) <-chan int {
	out := make(chan int)
	go func() {
		n := randInt(10000)
		for i := n; i > n-total*2; i-- {
			if i%2 != 0 {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

// HexIntGen generates random hex numbers
func HexIntGen(total int) <-chan int {
	out := make(chan int)
	hex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	go func() {
		for i := 0; i < total; i++ {
			n := randInt(len(hex) - 1)
			bigN := new(big.Int)
			i, wasSet := bigN.SetString(strings.Repeat(hex[n], n+1), 16)
			if wasSet {
				out <- int(i.Int64())
			}
		}
		close(out)
	}()
	return out
}

func randInt(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n)
}
