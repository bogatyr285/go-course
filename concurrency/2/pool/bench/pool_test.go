package benc_test

import (
	"bytes"
	"sync"
	"testing"
)

// Benchmark without pool
func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		buf.WriteString("Test String")
	}
}

// Benchmark with pool
var bufferPool = &sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Reset()
		buf.WriteString("Test String")
		bufferPool.Put(buf)
	}
}
