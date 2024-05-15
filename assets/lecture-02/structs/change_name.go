package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	person = &Person{
		Name: "Alice",
	}
}

func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
}

type Person struct {
	Name string
}

func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
	changeName2(person)
	fmt.Println(person.Name)
	changeName3(person)
	fmt.Println(person.Name)
}

func changeName(person *Person) *Person {
	return &Person{
		Name: "Alice0000",
	}
}

func changeName2(person *Person) {
	person.Name = "Alice2222"
}

func changeName3(person *Person) {
	*person = Person{
		Name: "qqq3333",
	}
}
