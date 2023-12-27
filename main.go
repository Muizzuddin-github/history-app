package api

import (
	"context"
	"crud/src"
	"crud/src/db"

	"github.com/gofiber/fiber/v2"
)



func Handler() *fiber.App{

	ctx := context.Background()
	db.CreateConn(ctx)
	defer db.CloseDB(ctx)

	server := src.Application()
	return server
}