package productService

import (
	"context"
	"errors"
	"time"

	"api.com/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type ProductServiceImpl struct {
	productcollection *mongo.Collection
	ctx            context.Context
}

func NewProductService(productcollection *mongo.Collection, ctx context.Context) *ProductServiceImpl {
	return &ProductServiceImpl {
		productcollection:productcollection, 
		ctx:ctx,
	}
}

func (u *ProductServiceImpl) CreateProduct(product *models.Product,userId *string) (string,error){

	if(product.Name == "" || product.Description == "" || product.Quantity == "0" || product.Type == ""){
		return "",errors.New("please fill all the fields")
	}

	if(*userId != product.UserId.Hex()){
		return "",errors.New("unauthorized")
	}

	product.Id = primitive.NewObjectID()
	product.CreatedAt = time.Now().Format("2006/01/02 15:04:05")
	result, err := u.productcollection.InsertOne(u.ctx, product)
	newId := result.InsertedID.(primitive.ObjectID).Hex()
	return newId,err
}

func (u *ProductServiceImpl) GetAllProducts() ([]*models.Product, error){
	var products []*models.Product
	query := bson.D{bson.E{Key: "listed", Value: true}}
	cursor, err := u.productcollection.Find(u.ctx, query)
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

func (u *ProductServiceImpl) GetProductById(productId *string) (*models.Product, error){
	var product *models.Product
	objID, _ := primitive.ObjectIDFromHex(*productId)
	query := bson.D{bson.E{Key: "_id", Value: objID}}
	err := u.productcollection.FindOne(u.ctx, query).Decode(&product)
	return product, err
}



