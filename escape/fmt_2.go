package main

import "fmt"

type Article struct {
	Title   string
	Content string
}

func main() {
	a := newArticle()
	println("main: a:", &a)
}

func newArticle() Article {
	a := Article{Title: "Нужен лишь простой советский..."}
	fmt.Println("func: a:", &a)
	return a
}
