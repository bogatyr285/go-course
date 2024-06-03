package main

import "fmt"

func main() {
	price := totalPrice(3, 10000, 132)
	fmt.Println(price)

	// price := totalPrice(3, 10000, 132)
	if price == 30396 {
		fmt.Println("function works")
	} else {
		fmt.Println("function is buggy")
	}
}

func totalPrice(nights, rate, cityTax uint) uint {
	return nights * (rate + cityTax)
}

// func main() {
// 	price := totalPrice(3, 10000, 132)
// 	if price == 30396 {
// 		fmt.Println("function works")
// 	} else {
// 		fmt.Println("function is buggy")
// 	}
// }
