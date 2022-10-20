package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	Area string `json:"area" bson:"area"`
	City string `json:"city" bson:"city"`
}

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	FirstName   string  `json:"firstName" bson:"firstName"`
	LastName    string  `json:"lastName" bson:"lastName"`
	PhoneNumber string     `json:"phoneNumber" bson:"phoneNumber"`
	Email       string  `json:"email" bson:"email"`
	Password    string  `json:"password" bson:"password"`
	Address     Address `json:"address" bson:"address"`
}

type UserLogin struct {
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	Password    string `json:"password" bson:"password"`
}
