package pg

import (
	"log"
	"main/pkg/api"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func UserLogin(c fiber.Ctx) error {
	var req api.LoginRequest
	var res api.LoginResult

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Msg":     "Invalid request body",
			"Success": false,
		})
	}

	if req.Email == "" || req.Password == "" {
		missingFields := []string{}
		if req.Email == "" {
			missingFields = append(missingFields, "Email")
		}
		if req.Password == "" {
			missingFields = append(missingFields, "Password")
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Msg":     "Missing required fields: " + strings.Join(missingFields, ", "),
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

		if err := db.QueryRow("select pas from users where email = $1;", req.Email).Scan(&res.Password); err != nil {
			log.Fatalf("Ошибка при валидации введенных данных %v", err)
		}
		var pas string = res.Password
		if req.Password != pas {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"Msg":     "Incorrect password",
				"Success": false,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Msg":     "Login successful",
			"Success": true,
		})

	}

	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"Msg":     "User with email " + req.Email + " not exists",
		"Success": false,
	})
}
