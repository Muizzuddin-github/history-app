package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Info struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	ImageUrl    string             `bson:"imageUrl" json:"imageUrl"`
	Created_at  string             `bson:"created_at" json:"created_at"`
}
