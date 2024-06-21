package main

type Article struct {
	Name  string
	views int
}

func main() {
	a := newArticle()
	println(a)
}
func newArticle() *Article {
	a := Article{Name: "А во время остановки в GC  происходит ...", views: 1}
	println(&a)
	return &a
}
