package main

import (
	"bytes"
	"io"
)

func print(w io.Writer, s string) {
	asBytes := []byte(s)
	w.Write(asBytes)
}

func main() {
	buf := bytes.Buffer{}
	print(&buf, "foobar")
}
