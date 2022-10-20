package tradeService

import (
	"context"
	"errors"
	"strconv"

	"api.com/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type TradeServiceImpl struct {
	tradecollection *mongo.Collection
	productcollection *mongo.Collection
	ctx            context.Context
}

func NewTradeService(tradecollection *mongo.Collection,productcollection *mongo.Collection, ctx context.Context) *TradeServiceImpl {
	return &TradeServiceImpl {
		tradecollection:tradecollection, 
		productcollection:productcollection,
		ctx:ctx,
	}
}

func (u *TradeServiceImpl) CreateTrade(trade *models.Trade,reqUserId *string) (string,error){

	var product *models.Product

	query := bson.D{bson.E{Key: "_id", Value: trade.ProductId1}}
	err := u.productcollection.FindOne(u.ctx, query).Decode(&product)

	if(err != nil){
		return "", errors.New("could not find product")
	}
	tradeQuantity, _ := strconv.Atoi(trade.Quantity1)
	productQuantity,_ := strconv.Atoi(product.Quantity)
	if(productQuantity < tradeQuantity || tradeQuantity <= 0){
		return "", errors.New("please check the quantity")
	}

	query = bson.D{bson.E{Key: "_id", Value: trade.ProductId2}}
	err = u.productcollection.FindOne(u.ctx, query).Decode(&product)

	if(err != nil){
		return "", errors.New("could not find product")
	}

	tradeQuantity, _ = strconv.Atoi(trade.Quantity2)
	productQuantity,_ = strconv.Atoi(product.Quantity)
	if(productQuantity < tradeQuantity && tradeQuantity <= 0){
		return "", errors.New("please check the quantity")
	}


	trade.Id = primitive.NewObjectID()
	trade.Status = "pending"
	result, err := u.tradecollection.InsertOne(u.ctx, trade)
	if err!=nil{
		return "",errors.New("could not create trade")
	}
	newId := result.InsertedID.(primitive.ObjectID).Hex()
	return newId,err
}

func (u *TradeServiceImpl) SetTradeStatus(tradeId *string,status string,reqUserId *string) error{
	var tradeModel *models.Trade
	objID, _ := primitive.ObjectIDFromHex(*tradeId)

	if(status != "cancelled" && status != "completed"){		
		return errors.New("invalid status")
	}

	query := bson.D{bson.E{Key: "_id", Value: objID}}
	err := u.tradecollection.FindOne(u.ctx, query).Decode(&tradeModel)

	if(err != nil){
		return errors.New("could not find trade")
	}


	if(tradeModel.TraderId.Hex() != *reqUserId){
		return errors.New("you are not authorized to update this trade")
	}

	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "status", Value: status}}}}
	result,_ := u.tradecollection.UpdateOne(u.ctx, query, update)
	if result.MatchedCount != 1{
		return errors.New("no matched document");
	}
	return nil
}

func (u *TradeServiceImpl) GetTrades(userId *string,status *string,reqUserId *string) ([]*models.Trade, error){
	var trades []*models.Trade
	objID, _ := primitive.ObjectIDFromHex(*userId)
	

	query := bson.M{
		"status": status,
		"$or": []interface{}{
			bson.M{"userId": objID},
			bson.M{"traderId": objID},
		},
	}

	// aggSearch := bson.M{"$match": query}

	// aggPopulate := bson.M{"$lookup": bson.M{
	// 	"from":         "Product",    
	// 	"localField":   "productId1", 
	// 	"foreignField": "_id",       
	// 	"as":           "product",    
	//   }}

	//   cursor, err := u.tradecollection.Aggregate(u.ctx, []bson.M{
	// 	aggPopulate,aggSearch,
	//   })

	cursor, err := u.tradecollection.Find(u.ctx,query)
	if err != nil{
		return nil, err
	}
	
	for cursor.Next(u.ctx){
		var trade models.Trade
		err := cursor.Decode(&trade)
		if err != nil{
			return nil, err
		}
		trades = append(trades, &trade)
	}
	if err := cursor.Err(); err != nil{
		return nil, err
	}
	
	cursor.Close(u.ctx)

	if(len(trades) > 0){
		if(trades[0].UserId.Hex() != *reqUserId && trades[0].TraderId.Hex() != *reqUserId){
			return nil, errors.New("unauthorized")
		}
		return trades, nil
	}else{
		return nil, errors.New("no trades found")
	}

}

func (u *TradeServiceImpl) CancellTrade(id *string,reqUserId *string) error{
	var tradeModel *models.Trade
	objID, _ := primitive.ObjectIDFromHex(*id)	

	query := bson.D{bson.E{Key: "_id", Value: objID}}
	err := u.tradecollection.FindOne(u.ctx, query).Decode(&tradeModel)

	if(err != nil){
		return errors.New("could not find trade")
	}

	if(tradeModel.UserId.Hex() != *reqUserId){
		return errors.New("you are not authorized to cancell this trade")
	}


	filter := bson.D{bson.E{Key: "_id", Value: objID}}
	result, err := u.tradecollection.DeleteOne(u.ctx, filter)
	if err != nil{
		return err
	}
	if result.DeletedCount != 1{
		return errors.New("no matched document");
	}
	return nil
	
}




