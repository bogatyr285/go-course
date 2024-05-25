package main

import "fmt"

// START OMIT

func main() {
	q := make([]string, 10)

	for i := range q {
		q[i] = "qq"
		fmt.Println(q[i])
	}
}

// END OMIT
