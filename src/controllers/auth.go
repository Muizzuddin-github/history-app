package controllers

import (
	"context"
	"crud/src/db"
	"crud/src/repository"
	"crud/src/requestbody"
	"crud/src/responsebody"
	"crud/src/utility"
	"crud/src/validation"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var Register fiber.Handler = func(c *fiber.Ctx) error {

	body := requestbody.Register{}
	c.BodyParser(&body)

	val := validation.Register(&body)
	if len(val) > 0{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(responsebody.Err{
			Errors: val,
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.Status(400)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}


	body.Password = string(hash)
	body.Created_at = utility.TimeNow()

	ctx, cancel := context.WithTimeout(context.Background(),time.Second * 10)
	defer cancel()

	usersCol := repository.NewAuthRepo(db.GetCollection("users"))
	err = usersCol.Register(ctx, &body)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(responsebody.Msg{
		Message: "Created",
	})
}

var Login fiber.Handler = func(c *fiber.Ctx) error {

	body := requestbody.Login{}
	c.BodyParser(&body)

	ctx := context.Background()
	userCol := repository.NewAuthRepo(db.GetCollection("users"))
	user, err := userCol.Login(ctx,&body)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(responsebody.Err{
			Errors: []string{"check your email or password"},
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))
	if err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(responsebody.Err{
			Errors: []string{"check your email or password"},
		})
	}

	token, err := utility.CreateToken(user.Id.Hex(),os.Getenv("SECRET_KEY"))
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}
	
	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: token,
		MaxAge: 24 * 60 * 60,
		HTTPOnly: false,
		Secure: true,
		SameSite: "none",
	})
	c.Status(fiber.StatusOK)
	return c.JSON(responsebody.Msg{
		Message: "success login",
	})
}

var Logout fiber.Handler = func(c *fiber.Ctx) error {
	c.ClearCookie("token")

	c.Status(fiber.StatusOK)
	return c.JSON(responsebody.Msg{
		Message: "logout success",
	})
}