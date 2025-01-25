package pg

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type registerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func UserRegister(c fiber.Ctx) error {
	var req registerRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Msg":     "Invalid request body",
			"Success": false,
		})
	}

	if req.Email == "" || req.Login == "" || req.Password == "" {
		missingFields := []string{}

		if req.Email == "" {
			missingFields = append(missingFields, "Email")
		}
		if req.Login == "" {
			missingFields = append(missingFields, "Login")
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
		log.Fatalf("Ошибка при выполнении SELECT запроса %v", err)
	}
	defer emailCheck.Close()

	if emailCheck.Next() {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"Msg":     "User with email " + req.Email + " already exists",
			"Success": false,
		})
	}

	_, err = db.Exec("INSERT INTO users (login, pas, email) VALUES ($1, $2, $3)", req.Login, req.Password, req.Email)
	if err != nil {
		log.Fatalf("Ошибка выполнения INSERT запроса: %v", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Msg":     "Новый пользователь зарегистрирован",
		"Success": true,
	})
}
