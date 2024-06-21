package main

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
	println("func: a:", &a)
	return a
}
