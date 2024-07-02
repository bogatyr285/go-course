package main

import "fmt"

// Node represents a single breadcrumb
type Node struct {
	path string
	next *Node
}

// LinkedList represents the breadcrumb chain
type LinkedList struct {
	head *Node
}

// AddBreadcrumb adds a new breadcrumb to the end of the list
func (list *LinkedList) AddBreadcrumb(path string) {
	newNode := &Node{path: path}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// RemoveLastBreadcrumb removes the last breadcrumb from the list
func (list *LinkedList) RemoveLastBreadcrumb() {
	if list.head == nil {
		return
	}
	if list.head.next == nil {
		list.head = nil
		return
	}
	current := list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

// DisplayBreadcrumbs displays the breadcrumb chain
func (list *LinkedList) DisplayBreadcrumbs() {
	if list.head == nil {
		fmt.Println("No breadcrumbs.")
		return
	}
	current := list.head
	for current != nil {
		fmt.Print(current.path)
		if current.next != nil {
			fmt.Print(" > ")
		}
		current = current.next
	}
	fmt.Println()
}

func main() {
	breadcrumbs := LinkedList{}
	breadcrumbs.AddBreadcrumb("Home")
	breadcrumbs.AddBreadcrumb("Products")
	breadcrumbs.AddBreadcrumb("Electronics")
	breadcrumbs.AddBreadcrumb("Laptops")

	// Display the initial breadcrumb trail
	breadcrumbs.DisplayBreadcrumbs()

	// Remove the last breadcrumb and display the updated trail
	breadcrumbs.RemoveLastBreadcrumb()
	breadcrumbs.DisplayBreadcrumbs()

	// Adding a new breadcrumb after removal
	breadcrumbs.AddBreadcrumb("Tablets")
	breadcrumbs.DisplayBreadcrumbs()
}
