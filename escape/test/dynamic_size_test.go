package test

import (
	"testing"
)

func StaticSize() {
	_ = make([]string, 0, 10)
	// println(a)
}

func DynamicSize(c int) {
	_ = make([]string, 0, c)
	// println(a)
}

func BenchmarkStaticSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StaticSize()
	}
}

func BenchmarkDynamicSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DynamicSize(10)
	}
}
