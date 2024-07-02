package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	odd := OddIntGen(10)
	c1 := FanOut(done, odd)
	c2 := FanOut(done, odd)
	c3 := FanOut(done, odd)
	c4 := FanOut(done, odd)

	display(c1, "c1")
	display(c2, "c2")
	display(c3, "c3")
	display(c4, "c4")
}

// FanOut считывает все значения из заданного входного канала и отправляет каждое значение в результирующий канал
// Нужен многократный вызов этой функции
// Создает несколько рутин для чтения из одного и того же канала, пока входной канал не будет закрыт
// Позволяет лучше распределить работу
func FanOut(done chan struct{}, input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range input {
			select {
			case <-done:
				return
			case out <- v:
			}
		}
		close(out)
	}()
	return out
}

func display(in <-chan int, label string) {
	for v := range in {
		fmt.Printf("value %s: %d\n", label, v)
	}
}

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

func randInt(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n)
}
