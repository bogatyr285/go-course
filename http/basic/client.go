package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	for i := 0; i < 10_000_000; i++ {
		go func() {
			resp, err := http.Get("http://localhost:8080/hello")
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()

			fmt.Println("Response status:", resp.Status)

			b, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			log.Println("body ", string(b))
		}()
	}
}
