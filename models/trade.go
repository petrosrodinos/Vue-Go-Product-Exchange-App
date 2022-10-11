package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trade struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	UserId          primitive.ObjectID `json:"userId" bson:"userId"`
	TraderId          primitive.ObjectID `json:"traderId" bson:"traderId"`
	ProductId1 primitive.ObjectID `json:"productId1" bson:"productId1"`
	ProductId2  primitive.ObjectID `json:"productId2" bson:"productId2"`
	Quantity1 string `json:"quantity1" bson:"quantity1"`
	Quantity2  string `json:"quantity2" bson:"quantity2"`
	Status string `json:"status" bson:"status"`
}

