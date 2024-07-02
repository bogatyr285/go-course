package main

import (
	"sort"
	"testing"
)

func binSearch(elements []int, element int) (bool, int) {
	i := sort.Search(len(elements), func(i int) bool {
		// log.Println(i)
		return elements[i] >= element
	})
	if i < len(elements) && elements[i] == element {
		// fmt.Printf("found element %d at index %d in %v\n", element, i, elements)
		return true, i
	}

	// fmt.Printf("element %d not found in %v\n", element, elements)
	return false, -1
}

func interpolationSearch(elements []int, element int) (bool, int) {
	var mid int
	var low int
	low = 0
	var high int
	high = len(elements) - 1
	for elements[low] < element && elements[high] > element {
		mid = low + ((element-elements[low])*(high-low))/(elements[high]-
			elements[low])
		if elements[mid] < element {
			low = mid + 1
		} else if elements[mid] > element {
			high = mid - 1
		} else {
			return true, mid
		}
	}
	if elements[low] == element {
		return true, low
	} else if elements[high] == element {
		return true, high
	}
	return false, -1
}

func BenchmarkInterpolationSearch(b *testing.B) {
	size := 1000
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = i
	}

	elementToFind := 500

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		interpolationSearch(elements, elementToFind)
	}
}

func BenchmarkBinSearch(b *testing.B) {
	// Sample data for benchmarking
	size := 1000
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = i
	}

	elementToFind := 500 // Change this value for different test cases

	// Reset the timer for a more accurate measurement
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		binSearch(elements, elementToFind)
	}
}
