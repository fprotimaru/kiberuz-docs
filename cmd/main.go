package main

import (
	"flag"
	"log"
	"net/http"

	"kiber-docs/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var port string

func main() {
	flag.StringVar(&port, "port", ":8001", "port")
	flag.Parse()

	app := fiber.New()
	app.Use("/docs", filesystem.New(filesystem.Config{
		Root:   http.FS(docs.FS),
		Index:  "index.html",
		Browse: true,
	}))
	log.Fatalln(app.Listen(port))
}
