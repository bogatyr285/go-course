package main

import (
	"io"
	"log"
)

func doSomething() error {
	var e *Err
	return e
}

// Err conforms to the error interface.
type Err string

// Error implements the error interface.
func (e *Err) Error() string {
	return string(*e)
}

func main() {
	var r io.Reader

	log.Println("r", r == nil)

	// err := doSomething()
	// if err != nil {
	// 	log.Fatalf("ERR %T %s", err, err)
	// }
	// log.Println("OK")
}
