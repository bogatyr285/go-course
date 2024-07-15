package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	if err := execute(); err != nil {
		log.Fatalln(err)
	}
}

func execute() (err error) {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(conn)

	request := []byte("GET / HTTP/1.1\r\n")
	for i := range request {
		time.Sleep(time.Second)
		part := request[i : i+1]
		log.Printf("write: %s", part)
		_, err = conn.Write(part)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	log.Println("request finished")

	return nil
}
