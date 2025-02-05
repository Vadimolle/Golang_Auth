package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func RegHandler(c fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	})
}

func LogHandler(c fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

func MainHandler(c fiber.Ctx) error {
	return c.Render("main", fiber.Map{
		"Title": "Home",
	})
}
