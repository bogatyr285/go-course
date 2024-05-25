package main

import (
	"fmt"
	"log"

	"io"

	"os"
)

// MultiWriter is a custom io.Writer that writes to multiple writers

type MultiWriter struct {
	writers []io.Writer
}

// Write writes data to all writers
func (mw *MultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range mw.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
	}
	return len(p), nil
}

func main() {
	// Open a file for logging
	file, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a MultiWriter for writing to both file and console
	multiWriter := &MultiWriter{
		writers: []io.Writer{file, os.Stdout, &TestStruct{}},
	}
	fmt.Fprintln(multiWriter, "This is a log message!")

	// You can use log package as well with MultiWriter
	logger := log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime)
	logger.Println("This is a log message with log package!")
}

type TestStruct struct {
}

func (s *TestStruct) Write(p []byte) (n int, err error) {
	return 0, nil
}
