package routers

import (
	"crud/src/controllers"
	"crud/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Info() *fiber.App {
	info := fiber.New()

	info.Post("/category/:idCategory",middlewares.Authorization,controllers.AddInfo)
	info.Delete("/category/:idCategory/info/:idInfo",middlewares.Authorization,controllers.DeleteInfo)
	info.Put("/category/:idCategory/info/:idInfo",middlewares.Authorization,controllers.UpdateInfoNoImage)
	info.Put("/category/:idCategory/info-image/:idInfo",middlewares.Authorization,controllers.UpdateInfoImage)


	return info
}