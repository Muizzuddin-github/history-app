package main

import (
	"context"
	"crud/src"
	"crud/src/db"
)



func main(){

	ctx := context.Background()
	db.CreateConn(ctx)
	defer db.CloseDB(ctx)

	server := src.Application()
	server.Listen("0.0.0.0:8080")
}