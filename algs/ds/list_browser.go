package main

import (
	"container/list"
	"fmt"
)

func main() {
	history := list.New()

	history.PushBack("google.com") //
	history.PushBack("golang.org") //
	history.PushBack("stackoverflow.com")

	// Current position in history
	current := history.Back()

	// Show current page
	fmt.Println("Currently Viewing:", current.Value)

	// Go back in history
	if current.Prev() != nil {
		current = current.Prev()
		fmt.Println("went back, currently viewing:", current.Value)
	} else {
		fmt.Println("no previous pages")
	}

	// Go forward in history
	if current.Next() != nil {
		current = current.Next()
		fmt.Println("went forward, currently viewing:", current.Value)
	} else {
		fmt.Println("no more pages")
	}
}
