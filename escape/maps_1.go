package main

type article struct {
	name  string
	views int
}

func main() {
	v := map[string]string{}
	qq(v)
	println(v)
}

func qq(m map[string]string) {
	m["qq"] = "changed"
}
