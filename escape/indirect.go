package main

type Article struct {
	Name  string
	views int
}

func main() {
	a := Article{
		Name:  "Биткоин купить надо в...",
		views: 99999,
	}
	incrementViews(&a)
	println("main views:", &a.views)
}

func incrementViews(a *Article) {
	a.views++
	println("func views:", &a.views)
}
