package env

import (
	"os"
	"testing"
)

// MYENV="wow3" go test -v course/tests/env
func TestFoo(t *testing.T) {
	env := os.Getenv("MYENV")
	t.Log("MYENV", env)
	//..
}
