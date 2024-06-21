package main

type article struct {
	name  string
	views int
}

func main() {
	a := newArticle()
	println("main: a:", &a)
}
func newArticle() article {
	a := Article{name: "Выучить с++ за 21 день может каждый. Нужно лишь ...", views: 1}
	println("func: a:", &a)
	return a
}
