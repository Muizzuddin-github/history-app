package repository

import (
	"context"
	"crud/src/requestbody"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type infoInterface interface {
	AddInfo(ctx context.Context, data *requestbody.Info, id string) (*mongo.UpdateResult, string, error)
	DeleteInfo(ctx context.Context, idCategory string, idInfo string) (*mongo.UpdateResult, error)
	UpdateNoImage(ctx context.Context, idCategory string, idInfo string, data *requestbody.InfoUpdateNoImage) (*mongo.UpdateResult, error)
	UpdateWithImage(ctx context.Context, idCategory string, idInfo string, url string) (*mongo.UpdateResult, error)
}

type infoRepo struct{
	Col *mongo.Collection
}


func (info *infoRepo) AddInfo(ctx context.Context, data *requestbody.Info, id string) (*mongo.UpdateResult,string,error){

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil, "", errors.New(err.Error())
	}

	genObjId := primitive.ObjectID.Hex(primitive.NewObjectID())

	filter := bson.M{"_id": bson.M{"$eq": objId}}
	addNestedDoc := bson.M{"$push" : bson.M{"info" : bson.M{
		"_id" : genObjId,
		"title" : data.Title,
		"description" : data.Description,
		"imageUrl" : data.Image,
		"created_at" : data.Created_at,

	}}}

	result,err := info.Col.UpdateOne(ctx,filter,addNestedDoc)
	if err != nil{
		return nil, "", errors.New(err.Error())
	}


	return result, id, nil
}

func(info *infoRepo) DeleteInfo(ctx context.Context, idCategory string, idInfo string) (*mongo.UpdateResult, error){
	objIdCategory, err := primitive.ObjectIDFromHex(idCategory)
	if err != nil{
		return nil, errors.New(err.Error())
	}

	filter := bson.M{"_id" : bson.M{"$eq" : objIdCategory}}
	update := bson.M{"$pull" : bson.M{"info" : bson.M{"_id" : bson.M{"$eq" : idInfo}}}}
	result, err := info.Col.UpdateOne(ctx,filter,update)
	if err != nil{
		return nil, errors.New(err.Error())
	}

	return result, nil
}

func(info *infoRepo) UpdateNoImage(ctx context.Context, idCategory string, idInfo string,data *requestbody.InfoUpdateNoImage) (*mongo.UpdateResult, error){
	objIdCategory, err := primitive.ObjectIDFromHex(idCategory)
	if err != nil{
		return nil, errors.New(err.Error())
	}


	filter := bson.M{"_id" : objIdCategory}
	update := bson.M{"$set" : bson.M{
		"info.$[items].title" : data.Title,
		"info.$[items].description" : data.Description,
		},
	}
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{
			bson.M{
				"items._id" : bson.M{
					"$eq" : idInfo,
				},
			},
		},
	}

	elementMatch := options.Update().SetArrayFilters(arrayFilters)

	result, err := info.Col.UpdateOne(ctx,filter,update,elementMatch)
	if err != nil{
		return nil, errors.New(err.Error())
	}

	return result,nil

}

func(info *infoRepo) UpdateWithImage(ctx context.Context, idCategory string, idInfo string,url string) (*mongo.UpdateResult, error){
	objIdCategory, err := primitive.ObjectIDFromHex(idCategory)
	if err != nil{
		return nil, errors.New(err.Error())
	}


	filter := bson.M{"_id" : objIdCategory}
	update := bson.M{"$set" : bson.M{
		"info.$[items].imageUrl" : url,
		},
	}
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{
			bson.M{
				"items._id" : bson.M{
					"$eq" : idInfo,
				},
			},
		},
	}

	elementMatch := options.Update().SetArrayFilters(arrayFilters)

	result, err := info.Col.UpdateOne(ctx,filter,update,elementMatch)
	if err != nil{
		return nil, errors.New(err.Error())
	}

	return result,nil

}


func NewInfoRepo(col *mongo.Collection) infoInterface{
	return &infoRepo{
		Col : col,
	}
}