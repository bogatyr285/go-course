// The server pattern is based on the simple idea of infinite loop
// that listens for connections and runs each connection in a go routine.
// Huge thanks to divan.dev. Check out the full resource here: https://divan.dev/posts/go_concurrency_visualize/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handler(c net.Conn) {
	reader := bufio.NewReader(c)
	body, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	// body, err := io.ReadAll(c)
	// if err != nil {
	// 	fmt.Println("r", err)
	// 	return
	// }
	fmt.Println("body", string(body))
	_, _ = c.Write([]byte("ok"))
	_ = c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		fmt.Println("waiting for request")
		c, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Println("run thread to process request: ", c.RemoteAddr())
		go handler(c)
	}
}
