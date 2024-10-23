package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		host, _ := os.Hostname()
		s := fmt.Sprintf("Hello, World! \nI'm: %s, ip: %s\n %s %s",
			host,
			GetLocalIP().String(),
			os.Getenv("HELM_RELEASE_NAME"),
			os.Getenv("HELM_RELEASE_NAMESPACE"),
		)
		return c.SendString(s)
	})

	app.Get("/crash", func(c *fiber.Ctx) error {
		log.Fatalln("intensional crash")
		return nil
	})

	app.Listen(":3000")
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
