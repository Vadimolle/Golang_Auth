package app

import (
	"log"
	"main/internal/pg"

	"github.com/gofiber/fiber/v3"
)

func Run() {
	//engine := html.New("./template", ".html")
	//app := fiber.New(fiber.Config{
	//	Views: engine,
	//})
	app := fiber.New()

	app.Post("/register", pg.UserRegister)
	app.Post("/login", pg.UserLogin)

	//app.Get("/", handlers.RegHandler)
	//app.Get("/login", handlers.LogHandler)
	//app.Get("/home", handlers.MainHandler)
	//app.Get("/*", static.New("./public"))

	log.Fatal(app.Listen(":3333"))
}
