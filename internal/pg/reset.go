package pg

import (
	"log"
	"main/pkg/api"

	"github.com/gofiber/fiber/v3"
)

func LoginReset(c fiber.Ctx) error {
	var req api.ResetRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Msg":     "Invalid request body",
			"Success": false,
		})
	}

	if req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Msg":     "Field email is missing",
			"Success": false,
		})
	}

	db, err := DBConn()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	emailCheck, err := db.Query("select email from users where email = $1;", req.Email)
	if err != nil {
		log.Fatalf("Ошибка при проверке на наличие email в базе %v", err)
	}
	defer emailCheck.Close()

	if emailCheck.Next() {

	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Msg":     "There is no user with this email",
		"Success": false,
	})
}
