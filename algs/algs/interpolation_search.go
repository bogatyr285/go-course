package main

import "fmt"

func InterpolationSearch(elements []int, element int) (bool, int) {
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

func main() {
	var elements []int
	elements = []int{2, 3, 5, 7, 9}
	var element int
	element = 7
	var found bool
	var index int
	found, index = InterpolationSearch(elements, element)
	fmt.Println(found, "found at", index)
}
