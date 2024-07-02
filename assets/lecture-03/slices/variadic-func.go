package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT

func printArgs(args ...string) {

}

func main() {
	q := q{}
	container
	b, _ := json.Marshal(q)

	fmt.Println(string(b))
}

type q struct {
	Text *string `json:"text,omitempty"`
}

func qq(v q) {

}
