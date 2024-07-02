package main

import (
	"math/rand"
	"sort"
	"testing"
)

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(1_000_000)
	}
	return slice
}

func quickSortBuiltin(arr []int) {
	sort.Ints(arr)
}

func quickSortStackOverflow(arr []int) {
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	pivotIndex := rand.Int() % len(arr)

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]

	quickSortStackOverflow(arr[:left])
	quickSortStackOverflow(arr[left+1:])
}

func BenchmarkQuickSortBuiltin(b *testing.B) {
	size := 10000
	arr := generateRandomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		quickSortBuiltin(arr)
		b.StopTimer()
	}
}

func BenchmarkQuickSortStackOverflow(b *testing.B) {
	size := 10000
	for i := 0; i < b.N; i++ {
		arr := generateRandomSlice(size)
		b.StartTimer()
		quickSortStackOverflow(arr)
		b.StopTimer()
	}
}
