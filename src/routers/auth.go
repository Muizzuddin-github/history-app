package routers

import (
	"crud/src/controllers"
	"crud/src/middlewares"

	"github.com/gofiber/fiber/v2"
)


func Auth() *fiber.App{

	auth := fiber.New()
	// auth.Post("/register",controllers.Register)
	auth.Post("/login",controllers.Login)
	auth.Post("/logout",controllers.Logout)
	auth.Get("/islogin",middlewares.Authorization,controllers.IsLogin)
	
	return auth
}