package test

import (
	"bytes"
	"io"
	"testing"
)

func writerInterface(w io.Writer, s string) {
	asBytes := []byte(s)
	w.Write(asBytes)
}

func writerBuffer(w *bytes.Buffer, s string) {
	asBytes := []byte(s)
	w.Write(asBytes)
}

func BenchmarkWriterInterface(b *testing.B) {
	buf := bytes.Buffer{}
	s := "foobar"
	for i := 0; i < b.N; i++ {
		writerInterface(&buf, s)
	}
}

func BenchmarkWriterBuffer(b *testing.B) {
	buf := bytes.Buffer{}
	s := "foobar"
	for i := 0; i < b.N; i++ {
		writerBuffer(&buf, s)
	}
}
