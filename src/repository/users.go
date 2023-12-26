package repository

import (
	"context"
	"crud/src/entity"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type usersRepoInterface interface{
	GetAllUser(ctx context.Context) ([]entity.Users, error)
	GetUserById(ctx context.Context, id string) (entity.Users,error)

}

type usersRepo struct {
	Col *mongo.Collection
}


func (user *usersRepo) GetAllUser(ctx context.Context) ([]entity.Users, error){
	cur, err := user.Col.Find(ctx,bson.D{})
	if err != nil{
		return nil, errors.New(err.Error())
	}
	defer cur.Close(ctx)

	users := []entity.Users{}

	for cur.Next(ctx){
		data := entity.Users{}
		err := cur.Decode(&data)
		if err != nil{
			return nil, errors.New(err.Error())
		}

		users = append(users, data)
	}

	return users,nil
}

func (user *usersRepo) GetUserById(ctx context.Context, id string) (entity.Users,error){
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return entity.Users{}, errors.New(err.Error())
	}

	filter := bson.M{"_id" : objId}

	result := user.Col.FindOne(ctx,filter)
	data := entity.Users{}
	result.Decode(&data)


	return data,nil
}



func NewUserRepo(col *mongo.Collection) usersRepoInterface {
	return &usersRepo{
		Col: col,
	}
}