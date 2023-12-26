package requestbody

import "go.mongodb.org/mongo-driver/bson/primitive"

type Info struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Title       string             `bson:"title" form:"title" validate:"required"`
	Description string             `bson:"description" form:"description" validate:"required"`
	ImageUrl    string             `bson:"imageUrl"`
	Created_at  string             `bson:"created_at"`
}

type InfoUpdateNoImage struct{
	Title       string             `bson:"title" form:"title" validate:"required"`
	Description string             `bson:"description" form:"description" validate:"required"`
}

type InfoUpdateWithImage struct{
	Title       string             `bson:"title" form:"title" validate:"required"`
	Description string             `bson:"description" form:"description" validate:"required"`
	ImageUrl    string             `bson:"imageUrl"`
}