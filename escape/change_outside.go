package main

type Container struct {
	value *int
}

func (c *Container) SetValue(v *int) {
	c.value = v
}

// func (c *Container) SetValueHeap(v *int) {
// 	c.value = v
// }

func main() {
	a := 1337
	container := &Container{}
	container.SetValue(&a)
	println(container.value)
}
