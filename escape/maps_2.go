package main

type article struct {
	name  string
	views int
}

func main() {
	v := map[string]*article{
		"a": &article{
			name:  "How to cook pizza...",
			views: 1,
		}}
	println(v)
}
