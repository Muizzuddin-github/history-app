package entity

import "go.mongodb.org/mongo-driver/bson/primitive"



type Category struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	Category string `bson:"category" json:"category"`
	Created_at string `bson:"created_at" json:"created_at"`
	Info []Info `bson:"info" json:"info"`
}