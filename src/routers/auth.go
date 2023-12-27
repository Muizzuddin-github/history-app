package routers

import (
	"crud/src/controllers"

	"github.com/gofiber/fiber/v2"
)


func Auth() *fiber.App{

	auth := fiber.New()
	// auth.Post("/register",controllers.Register)
	auth.Post("/login",controllers.Login)
	auth.Post("/logout",controllers.Logout)
	
	return auth
}