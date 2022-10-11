package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExchangeFor struct {
	Name string `json:"name" bson:"name"`
	Quantity string `json:"quantity" bson:"quantity"`
}

type Product struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	UserId          primitive.ObjectID `json:"userId" bson:"userId"`
	Name   string  `json:"name" bson:"name"`
	Description    string  `json:"description" bson:"description"`
	Quantity string     `json:"quantity" bson:"quantity"`
	Image       string  `json:"image" bson:"image"`
	Type    string  `json:"type" bson:"type"`
	Listed bool `json:"listed" bson:"listed"`
	ExchangeFor     ExchangeFor `json:"exchangeFor" bson:"exchangeFor"`
	CreatedAt string `json:"createdAt" bson:"createdAt"`
}

