package main

import (
	"bytes"
	"reflect"
)

func doPrint(b *bytes.Buffer, a []any) {
	prevString := false
	for argNum, arg := range a {
		isString := arg != nil && reflect.TypeOf(arg).Kind() == reflect.String
		// Add a space between two non-string arguments.
		if argNum > 0 && !isString && !prevString {
			b.WriteByte(' ')
		}
		prevString = isString
	}
}

func main() {
	w := bytes.Buffer{}
	doPrint(&w, []any{"foobar"})
}
