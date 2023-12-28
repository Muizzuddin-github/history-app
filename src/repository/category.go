package repository

import (
	"context"
	"crud/src/entity"
	"crud/src/requestbody"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type categoryInterface interface {
	AddCategory(ctx context.Context, data *requestbody.AddCategory) (string,error)
	GetCategory(ctx context.Context) ([]entity.Category, error)
	UpdateCategory(ctx context.Context, data *requestbody.UpdateCategory, id string) (*mongo.UpdateResult, error)
	DeleteCategory(ctx context.Context, id string) (*mongo.DeleteResult, error)
}


type categoryRepo struct{
	Col *mongo.Collection
}

func (category *categoryRepo) AddCategory(ctx context.Context, data *requestbody.AddCategory) (string,error){

	result, err := category.Col.InsertOne(ctx,data)
	if err != nil{
		return "", errors.New(err.Error())
	}

	insertId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok{
		return "", errors.New(err.Error())
	}

	return insertId.Hex(),nil
}


func (category *categoryRepo) GetCategory(ctx context.Context) ([]entity.Category, error){
	cur, err := category.Col.Find(ctx,bson.D{})
	if err != nil{
		return nil, errors.New(err.Error())
	}

	categories := []entity.Category{}

	for cur.Next(ctx){

		cat := entity.Category{}
		err := cur.Decode(&cat)
		if err != nil{
			return nil, errors.New(err.Error())
		}

		categories = append(categories, cat)
	}

	return categories,nil
}

func (category *categoryRepo) UpdateCategory(ctx context.Context, data *requestbody.UpdateCategory,id string) (*mongo.UpdateResult, error){

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil, errors.New(err.Error())
	}
	filter := bson.M{"_id": bson.M{"$eq": objId}}
	update := bson.M{"$set" : bson.M{
		"category" : data.Category,
	}}
	result, err := category.Col.UpdateOne(ctx,filter,update)
	if err != nil{
		return nil, errors.New(err.Error())
	}
	return result, nil
}


func (category *categoryRepo) DeleteCategory(ctx context.Context, id string) (*mongo.DeleteResult, error){

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil, errors.New(err.Error())
	}

	filter := bson.M{"_id": bson.M{"$eq": objId}}
	result, err := category.Col.DeleteOne(ctx,filter)
	if err != nil{
		return nil,errors.New(err.Error())
	}

	return result, nil
}


func NewCategoryRepo(col * mongo.Collection) categoryInterface{
	return &categoryRepo{
		Col: col,
	}
}