package src

import (
	"crud/src/routers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func Application() *fiber.App {

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		IdleTimeout: time.Second * 5,
		ReadTimeout: time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	

	app.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("index endpoint")
	})

	app.Mount("/",routers.Auth())
	// app.Mount("/",routers.Users())
	app.Mount("/",routers.Category())
	app.Mount("/",routers.Info())

	return app
}