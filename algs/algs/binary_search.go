package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	// var elements []int
	// elements = []int{1, 3, 16, 10, 45, 31, 28, 36, 45, 75}
	// var element int
	// element = 36
	// binSearch(elements, element)

	// GuessingGame()
}

func binSearch(elements []int, element int) (bool, int) {
	var i int
	i = sort.Search(len(elements), func(i int) bool {
		log.Println(i)
		return elements[i] >= element
	})
	if i < len(elements) && elements[i] == element {
		fmt.Printf("found element %d at index %d in %v\n", element, i, elements)
		return true, i
	}

	fmt.Printf("element %d not found in %v\n", element, elements)
	return false, -1
}
