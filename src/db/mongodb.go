package db

import (
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func CreateConn(ctx context.Context){
	opt := options.Client().ApplyURI(os.Getenv("DB"))
	client, err := mongo.Connect(ctx,opt)
	if err != nil{
		panic(err.Error())
	}

	err = client.Ping(ctx,nil)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println("database kenek")

	db = client.Database("informan")
}

func GetCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func CloseDB(ctx context.Context){
	db.Client().Disconnect(ctx)
}