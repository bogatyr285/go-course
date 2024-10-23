package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if err := execute(); err != nil {
		os.Exit(1)
	}
}

func execute() (err error) {
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
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

	r := bufio.NewReader(conn)
	const delim = '\n'
	line, err := r.ReadString(delim)
	if err != nil {
		if err != io.EOF {
			log.Println(err)
		}
		log.Printf("received: %s\n", line)
		return
	}
	log.Printf("received: %s\n", line)

	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		log.Printf("invalid request line: %s", line)
		return
	}

	path := parts[1]

	switch path {
	case "/":
		err = writeIndex(conn)
	case "/operations.csv":
		err = writeOperations(conn)
	default:
		err = write404(conn)
	}
	if err != nil {
		log.Println(err)
		return
	}
}

func writeIndex(writer io.Writer) error {
	username := "Василий"
	balance := "1 337.0" // TODO: написать функцию-форматтер

	page, err := os.ReadFile("web/template/index.html")
	if err != nil {
		return err
	}
	page = bytes.ReplaceAll(page, []byte("{username}"), []byte(username))
	page = bytes.ReplaceAll(page, []byte("{balance}"), []byte(balance))

	return writeResponse(writer, 200, []string{
		"Content-Type: text/html;charset=utf-8",
		fmt.Sprintf("Content-Length: %d", len(page)),
		"Connection: close",
	}, page)
}

func writeOperations(writer io.Writer) error {
	// TODO: Generate CSV
	page := []byte("xxxx,0001,0002,1592373247\n")

	return writeResponse(writer, 200, []string{
		"Content-Type: text/csv",
		fmt.Sprintf("Content-Length: %d", len(page)),
		"Connection: close",
	}, page)
}

func write404(writer io.Writer) error {
	page, err := os.ReadFile("web/template/404.html")
	if err != nil {
		return err
	}

	return writeResponse(writer, 200, []string{
		"Content-Type: text/html;charset=utf-8",
		fmt.Sprintf("Content-Length: %d", len(page)),
		"Connection: close",
	}, page)
}

func writeResponse(
	writer io.Writer,
	status int,
	headers []string,
	content []byte,
) error {
	const CRLF = "\r\n"
	var err error

	w := bufio.NewWriter(writer)
	_, err = w.WriteString(fmt.Sprintf("HTTP/1.1 %d OK%s", status, CRLF))
	if err != nil {
		return err
	}

	for _, h := range headers {
		_, err = w.WriteString(h + CRLF)
		if err != nil {
			return err
		}
	}

	_, err = w.WriteString(CRLF)
	if err != nil {
		return err
	}
	_, err = w.Write(content)
	if err != nil {
		return err
	}

	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}
