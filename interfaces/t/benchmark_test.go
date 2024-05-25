package t

import "testing"

type Doer interface{ Do(a, b int32) int32 }

type DoerStruct struct{ id int32 }

func (adder DoerStruct) Do(a, b int32) int32 { return a + b }

func BenchmarkDirect(b *testing.B) {
	adder := DoerStruct{id: 1337}
	for i := 0; i < b.N; i++ {
		adder.Do(10000000, 1)
	}
}

func BenchmarkInterface(b *testing.B) {
	adder := Doer(DoerStruct{id: 1337})
	for i := 0; i < b.N; i++ {
		adder.Do(10000000, 1)
	}
}
