package tradeController

import (
	"net/http"

	"api.com/api/models"
	"api.com/api/services/tradeService"
	"github.com/gin-gonic/gin"
)


type TradeController struct {
	TradeService tradeService.TradeService
}

func New(tradeService tradeService.TradeService) TradeController{
	return TradeController{
		TradeService: tradeService,
	}
}

func (uc *TradeController) CreateTrade(ctx *gin.Context){
	var trade models.Trade
	if err := ctx.ShouldBindJSON(&trade); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message": err.Error()})
		return
	}
	reqUserId,_ := ctx.Get("userId")
	finalUserId := reqUserId.(string)
	_,err := uc.TradeService.CreateTrade(&trade,&finalUserId)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Trade created successfully"})
}

func (uc *TradeController) SetTradeStatus(ctx *gin.Context) {
	tradeId := ctx.Param("id")
	status := ctx.Param("status")
	reqUserId,_ := ctx.Get("userId")
	finalUserId := reqUserId.(string)
   err := uc.TradeService.SetTradeStatus(&tradeId,status,&finalUserId)
   if err != nil{
	   ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
	   return
   }
   ctx.JSON(http.StatusOK,gin.H{"message": "trade status updated successfully"})
}

func (uc *TradeController) GetTrades(ctx *gin.Context){
	status := ctx.Param("status")
	userId := ctx.Param("userId")
	reqUserId,_ := ctx.Get("userId")
	finalUserId := reqUserId.(string)
	trades,err := uc.TradeService.GetTrades(&userId,&status,&finalUserId)
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Products found successfully", "trades": trades})

}

func (uc *TradeController) CancellTrade(ctx *gin.Context){
	id := ctx.Param("id")
	reqUserId,_ := ctx.Get("userId")
	finalUserId := reqUserId.(string)
	err := uc.TradeService.CancellTrade(&id,&finalUserId)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "trade cancelled successfully"})
}



