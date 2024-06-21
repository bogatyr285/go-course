package main

func f(a interface{}) {
	switch a.(type) {
	case int:
		println("int")
	case []int:
		println("[]int")
	case *int:
		println("*int")
	case *[]int:
		println("*[]int")
	}
}

func main() {
	var x = []int{0, 1, 2}

	f(x)
	f(x[0])
	f(&x)
	f(&x[0])
}
