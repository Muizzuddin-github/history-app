package repository

import (
	"context"
	"crud/src/entity"
	"crud/src/requestbody"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authInterface interface {
	Register(ctx context.Context, data *requestbody.Register) error
	Login(ctx context.Context, data *requestbody.Login) (entity.Users, error)
}


type auth struct{
	Col *mongo.Collection
}


func (auth *auth) Register(ctx context.Context, data *requestbody.Register) error{
	_, err := auth.Col.InsertOne(ctx,data)
	if err != nil{
		return errors.New(err.Error())
	}

	return nil
}

func(auth *auth) Login(ctx context.Context, data *requestbody.Login) (entity.Users, error){
	result := auth.Col.FindOne(ctx,bson.D{{Key: "email",Value: data.Email}})

	user := entity.Users{}
	err := result.Decode(&user)

	if err == mongo.ErrNoDocuments{
		return entity.Users{}, errors.New(err.Error())
	}

	return user, nil

}


func NewAuthRepo(col *mongo.Collection) authInterface{
	return &auth{
		Col: col,
	}
}