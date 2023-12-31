package routers

import (
	"crud/src/controllers"
	"crud/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Category() *fiber.App{

	category := fiber.New()

	category.Get("/category", middlewares.Authorization, controllers.GetAllCategory)
	category.Post("/category",middlewares.Authorization,controllers.AddCategory)
	category.Put("/category/:id",middlewares.Authorization,controllers.UpdateCategory)
	category.Delete("/category/:id",middlewares.Authorization,controllers.DeleteCategory)

	return category
}