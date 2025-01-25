package main

import (
	"log"
	"main/handlers"
	"main/pg"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./template", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Post("/register", pg.UserRegister)
	app.Post("/login", pg.UserLogin)

	app.Get("/", handlers.RegHandler)
	app.Get("/login", handlers.LogHandler)
	app.Get("/home", handlers.MainHandler)
	app.Get("/*", static.New("./public"))

	log.Fatal(app.Listen(":3333"))
}
