package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if err := execute(); err != nil {
		os.Exit(1)
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

	_, err = conn.Write([]byte("hello from client!\r\n"))
	if err != nil {
		log.Println(err)
		return err
	}

	r := bufio.NewReader(conn)
	const delim = '\n'
	line, err := r.ReadString(delim)
	if err != nil {
		return err
	}
	log.Println("l", line)

	return nil
}
