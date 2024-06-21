package main

type Article struct {
	name  string
	views int
}

func main() {
	a := Article{name: "Выучить с++ за 21 день может каждый. Нужно лишь ...", views: 1}
	println("func: a:", &a)
	println("main: a:", &a)
}
