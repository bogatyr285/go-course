package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	if err := execute(); err != nil {
		log.Fatalln(err)
	}
}

func execute() (err error) {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if cerr := listener.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}()
	for {
		conn, err := listener.Accept() // для клиентов
		if err != nil {
			log.Println(err)
			continue
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			log.Println(cerr)
		}
	}()

	reader := bufio.NewReader(conn)
	const delim = '\n'
	line, err := reader.ReadString(delim)
	if err != nil {
		if err != io.EOF {
			log.Println(err)
		}
		log.Printf("received: %s\n", line)
		return
	}
	log.Printf("received: %s\n", line)

	writer := bufio.NewWriter(conn)
	_, err = writer.WriteString("hello from server!\r\n")
	if err != nil {
		log.Println(err)
		return
	}
	err = writer.Flush()
	if err != nil {
		log.Println(err)
		return
	}

	for {
		line, err := reader.ReadString(delim)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				break
			}
			log.Printf("received loop: %s\n", line)
			return
		}
		log.Printf("received loop: %s\n", line)
	}
}
