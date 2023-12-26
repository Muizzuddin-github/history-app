package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"` 
	Username   string `bson:"username" json:"username"`
	Email      string `bson:"email" json:"email"`
	Password   string `bson:"password" json:"password"`
	Created_at string `bson:"created_at" json:"created_at"`
}