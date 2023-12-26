package middlewares

import (
	"context"
	"crud/src/db"
	"crud/src/repository"
	"crud/src/responsebody"
	"crud/src/utility"
	"os"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

var Authorization fiber.Handler = func(c *fiber.Ctx) error {

	cookie := string(c.Request().Header.Cookie("token"))
	if cookie == ""{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(responsebody.Err{
			Errors: []string{"Unauthorized"},
		})
	}

	result, err := utility.DecodeToken(cookie,os.Getenv("SECRET_KEY"))
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}


	idUser, ok := (*result)["id"]
	if !ok{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	ctx := context.Background()
	usersCol := repository.NewUserRepo(db.GetCollection("users"))
	user, err := usersCol.GetUserById(ctx, idUser.(string))
	if err != nil{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	keys := reflect.ValueOf(user)

	if keys.IsZero(){
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	return c.Next()
}