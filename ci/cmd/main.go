package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
)

var GlobalVar = make([]string, 10)

func main() {
	var x = 0
	_ = x

	// Inefficient assignment in loop (staticcheck )
	slice := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		slice = append(slice, i)
	}

	// Inefficient string concatenation (staticcheck)
	str := "Hello"
	str += " World"
	fmt.Println(str)

	// global variable
	GlobalVar[0] = "Set Global Var"

	// Hardcoded credentials
	username := "admin"
	password := "P@ssw0rd!"

	connStr := fmt.Sprintf("user=%s password=%s dbname=mydb sslmode=disable", username, password)
	log.Println(connStr)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}
	_, err := client.Get("https://golang.org/")
	if err != nil {
		fmt.Println(err)
	}

	h := http.Header{}
	h["etag"] = []string{"1234"}
	h.Add("etag", "5678")
	fmt.Println(h)
}

// Some uncalled function (staticcheck will catch this)
func uncalledFunc() {
	fmt.Println("This function is never called")
}

// func main() {
// 	fmt.Println("qqq")
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		host, _ := os.Hostname()
// 		s := fmt.Sprintf("Hello, World!\nI'm: %s, ip: %s", host, GetLocalIP().String())
// 		return c.SendString(s)
// 	})

// 	app.Get("/crash", func(c *fiber.Ctx) error {
// 		log.Fatalln("intensional crash")
// 		return nil
// 	})

// 	err := app.Listen(":3000")
// 	if err != nil {
// 		log.Fatalln("listen err", err)
// 	}
// }

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
