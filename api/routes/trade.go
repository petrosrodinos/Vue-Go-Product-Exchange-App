package routes

import (
	"api.com/api/controllers/tradeController"
	"api.com/api/middlewares"
	"github.com/gin-gonic/gin"
)

type TradeRoutes struct {
	tradecontroller tradeController.TradeController
	// mediacontroller mediaController.mediaController
}

func NewTradeRoutes(tradecontroller tradeController.TradeController) TradeRoutes{
	return TradeRoutes{
		tradecontroller: tradecontroller,
	}
}

func (uc *TradeRoutes) RegisterTradeRoutes(rg *gin.RouterGroup) {
	traderoutes := rg.Group("/trade",middlewares.AuthorizeJWT())
	traderoutes.POST("", uc.tradecontroller.CreateTrade)
	traderoutes.GET("/:userId/:status", uc.tradecontroller.GetTrades)
	traderoutes.PUT("/:id/:status", uc.tradecontroller.SetTradeStatus)
	traderoutes.DELETE("/:id", uc.tradecontroller.CancellTrade)

}