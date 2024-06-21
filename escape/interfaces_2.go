package main

import (
	"bytes"
)

func print(w *bytes.Buffer, s string) {
	asBytes := []byte(s)
	w.Write(asBytes)
}

func main() {
	buf := bytes.Buffer{}
	print(&buf, "foobar")
}
