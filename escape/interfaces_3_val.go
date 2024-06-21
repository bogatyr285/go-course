package main

type Writer interface {
	Write(b []byte) (int, error)
}

type writer struct{}

func (w writer) Write(b []byte) (int, error) {
	return 0, nil
}

func print(w Writer) {
	w.Write([]byte("foobar"))
}

func main() {
	var w Writer = writer{}
	print(w)
	// print(writer{})
}
