package routers

import (
	"crud/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Users() *fiber.App{
	users := fiber.New()

	users.Get("/users",controllers.HandleAllUser)

	return users
}