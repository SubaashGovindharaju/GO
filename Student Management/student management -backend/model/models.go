package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentManage struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string  `json:"Name,omitempty" bson:"Name,omitempty"`
	Gender string  `json:"Gender,omitempty" bson:"Gender,omitempty"`
	Age string `json:"Age,omitempty" bson:"Age,omitempty"`
}

