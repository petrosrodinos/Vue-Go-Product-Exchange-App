package userService

import (
	"context"
	"errors"

	"api.com/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type UserServiceImpl struct {
	productcollection *mongo.Collection
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(productcollection *mongo.Collection,usercollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl {
		productcollection:productcollection,
		usercollection:usercollection, 
		ctx:ctx,
	}
}

func (u *UserServiceImpl) GetProductsByUserId(userId *string) ([]*models.Product, error){
	var products []*models.Product
	objID, _ := primitive.ObjectIDFromHex(*userId)
	query := bson.D{bson.E{Key: "userId", Value: objID}}
	cursor, err := u.productcollection.Find(u.ctx,query)
	if err != nil{
		return nil, err
	}
	for cursor.Next(u.ctx){
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil{
			return nil, err
		}
		products = append(products, &product)
	}
	if err := cursor.Err(); err != nil{
		return nil, err
	}
	
	cursor.Close(u.ctx)

	if len(products) == 0{
		return nil, errors.New("no products found")
	}
	return products, nil
}

func (u *UserServiceImpl) GetUserById(userId *string) (*models.User, error){
	var user *models.User
	objID, _ := primitive.ObjectIDFromHex(*userId)
	query := bson.D{bson.E{Key: "_id", Value: objID}}
	opts := options.FindOne().SetProjection(bson.M{"password": 0})
	err := u.usercollection.FindOne(u.ctx,
		query,
		opts,
	).Decode(&user)

	return user, err
}

func (u *UserServiceImpl) GetListedProductsByUserId(userId *string) ([]*models.Product, error){
	var products []*models.Product
	objID, _ := primitive.ObjectIDFromHex(*userId)
	query := bson.D{bson.E{Key: "userId", Value: objID}, bson.E{Key: "listed", Value: true}}
	cursor, err := u.productcollection.Find(u.ctx,query)
	if err != nil{
		return nil, err
	}
	for cursor.Next(u.ctx){
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil{
			return nil, err
		}
		products = append(products, &product)
	}
	if err := cursor.Err(); err != nil{
		return nil, err
	}
	
	cursor.Close(u.ctx)

	if len(products) == 0{
		return nil, errors.New("no products found")
	}
	return products, nil
}

func (u *UserServiceImpl) GetUnlistedProductsByUserId(userId *string) ([]*models.Product, error){
	var products []*models.Product
	objID, _ := primitive.ObjectIDFromHex(*userId)
	query := bson.D{bson.E{Key: "userId", Value: objID}, bson.E{Key: "listed", Value: false}}
	cursor, err := u.productcollection.Find(u.ctx,query)
	if err != nil{
		return nil, err
	}
	for cursor.Next(u.ctx){
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil{
			return nil, err
		}
		products = append(products, &product)
	}
	if err := cursor.Err(); err != nil{
		return nil, err
	}
	
	cursor.Close(u.ctx)

	if len(products) == 0{
		return nil, errors.New("no products found")
	}
	return products, nil
}






