package authService

import (
	"context"
	"errors"
	"log"

	"api.com/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func getHash(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}

func NewAuthService(usercollection *mongo.Collection, ctx context.Context) *AuthServiceImpl {
	return &AuthServiceImpl {
		usercollection:usercollection, 
		ctx:ctx,
	}
}

func (u *AuthServiceImpl) RegisterUser(user *models.User) (string,error){
	query := bson.D{bson.E{Key: "phoneNumber", Value: user.PhoneNumber}}
	error := u.usercollection.FindOne(u.ctx, query).Decode(&user)

	if error == nil{
		return "",errors.New("user already exists")
	}

	user.Password = getHash([]byte(user.Password))
	user.Id = primitive.NewObjectID()
	result, err := u.usercollection.InsertOne(u.ctx, user)
	newId := result.InsertedID.(primitive.ObjectID).Hex()
	return newId,err
}

func (u *AuthServiceImpl) LoginUser(user *models.UserLogin) (string,error){
	var dbUser *models.User

	err:= u.usercollection.FindOne(u.ctx, bson.M{"phoneNumber":user.PhoneNumber}).Decode(&dbUser)

	if err!=nil{
		return "",errors.New("invalid credentials")
	}

	userPass:= []byte(user.Password)
  	dbPass:= []byte(dbUser.Password)

	passErr:= bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil{
		return "",errors.New("invalid credentials")
	}	

	userId := dbUser.Id.Hex()

	return userId,nil

}

func (u *AuthServiceImpl) GetUser(name *string) (*models.User, error){
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *AuthServiceImpl) GetAll() ([]*models.User, error){
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{})
	if err != nil{
		return nil, err
	}
	for cursor.Next(u.ctx){
		var user models.User
		err := cursor.Decode(&user)
		if err != nil{
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil{
		return nil, err
	}
	
	cursor.Close(u.ctx)

	if len(users) == 0{
		return nil, errors.New("no matched document")
	}
	return users, nil
}

func (u *AuthServiceImpl) UpdateUser(user *models.User) error{
	// filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	// update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.firstName},bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}
	// result,_ := u.usercollection.UpdateOne(u.ctx, filter, update)
	// if result.MatchedCount != 1{
	// 	return errors.New("no matched document");
	// }
	return nil
}

func (u *AuthServiceImpl) DeleteUser(name *string) error{
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, err := u.usercollection.DeleteOne(u.ctx, filter)
	if err != nil{
		return err
	}
	if result.DeletedCount != 1{
		return errors.New("no matched document");
	}
	return nil
	
}