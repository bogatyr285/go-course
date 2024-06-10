package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	c := displayData(prepareData(generateData()))
	for data := range c {
		log.Printf("Items: %+v", data)
	}
}

// tp access in different files
type Operation struct {
	id       int64
	multiply int64
	addition int64
}

func prepareData(ic <-chan int64) <-chan Operation {
	oc := make(chan Operation)
	go func() {
		for id := range ic {
			input := Operation{id: id, multiply: id * 2, addition: id + 5}
			oc <- input
		}
		close(oc)
	}()
	return oc
}

func generateData() <-chan int64 {
	c := make(chan int64)
	const filePath = "integer.txt"
	go func() {
		file, _ := os.Open(filePath)
		defer close(c)
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")
			integer, _ := strconv.ParseInt(line, 10, 0)
			c <- integer

			if err == io.EOF {
				break
			}
		}
	}()

	return c
}

func displayData(ic <-chan Operation) <-chan string {
	oc := make(chan string)

	go func() {
		wg := &sync.WaitGroup{}
		for input := range ic {
			wg.Add(1)
			go concatenateValue(input, oc, wg)
		}
		wg.Wait()
		close(oc)
	}()
	return oc
}

func concatenateValue(input Operation, output chan<- string, wg *sync.WaitGroup) {
	concatenated := fmt.Sprintf("id: %d, multiplication: %d, addition %d", input.id, input.multiply, input.addition)
	output <- concatenated
	wg.Done()
}
