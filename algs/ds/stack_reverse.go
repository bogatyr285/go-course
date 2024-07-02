package main

import (
	"fmt"
)

type Stack struct {
	elements []rune
}

func (s *Stack) Push(char rune) {
	s.elements = append(s.elements, char)
}

// Pop removes and returns the last added element from the stack
func (s *Stack) Pop() (rune, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return element, nil
}

// ReverseString reverses the input string using a stack
func ReverseString(input string) (string, error) {
	var stack Stack

	// Push all characters of the string onto the stack
	for _, char := range input {
		stack.Push(char)
	}

	// Pop all characters from the stack to form the reversed string
	var reversed []rune
	for range input {
		char, err := stack.Pop()
		if err != nil {
			return "", err
		}
		reversed = append(reversed, char)
	}

	return string(reversed), nil
}

func main() {
	input := "Hello, Web Development!"
	reversed, err := ReverseString(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Original String: %s\n", input)
	fmt.Printf("Reversed String: %s\n", reversed)
}
