package main

import (
	"context"
	"crud/src"
	"crud/src/db"
)


func main() {
	// mongodb connection
	ctx := context.Background()
	db.CreateConn(ctx)
	defer db.CloseDB(ctx)

	server := src.Application()
	server.Listen(":8080")
}