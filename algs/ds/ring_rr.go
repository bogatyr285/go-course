package main

import (
	"container/ring"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// List of backend servers
	backends := []string{
		"http://localhost:8001",
		"http://localhost:8002",
		"http://localhost:8003",
	}

	// Create a ring with the size equal to number of backends
	r := ring.New(len(backends))

	for _, backend := range backends {
		r.Value = backend
		r = r.Next()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		backend := r.Value.(string)
		r = r.Next()

		// Perform a reverse proxy to the selected backend
		url, _ := url.Parse(backend)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, req)
	})

	fmt.Println("Load balancer is running at :8080")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalln(err)
	}
}
