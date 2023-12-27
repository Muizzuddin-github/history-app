package main

import (
	"context"
	"crud/src"
	"crud/src/db"
	"os"
)



func main(){

	ctx := context.Background()
	db.CreateConn(ctx)
	defer db.CloseDB(ctx)

	port := os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}

	server := src.Application()
	server.Listen(":" + port)
}