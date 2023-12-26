package controllers

import (
	"context"
	"crud/src/db"
	"crud/src/repository"
	"crud/src/responsebody"

	"github.com/gofiber/fiber/v2"
)

var HandleAllUser fiber.Handler = func(c *fiber.Ctx) error {

	ctx := context.Background()
	usersCol := repository.NewUserRepo(db.GetCollection("users"))
	users, err := usersCol.GetAllUser(ctx)
	if err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(responsebody.Data{
		Message: "Success",
		Data: users,
	})
}

