package requestbody

import "crud/src/entity"

type AddCategory struct {
	Category   string `json:"category" validate:"required"`
	Created_at string `json:"created_at"`
	Info       []entity.Info `json:"info"`
}

type UpdateCategory struct{
	Category string `bson:"category" json:"category" validate:"required"`
}