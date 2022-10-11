package tradeService

import "api.com/api/models"

type TradeService interface {
	CreateTrade(*models.Trade,*string) (string,error)
	SetTradeStatus(*string,string,*string) error
	GetTrades(*string,*string,*string) ([]*models.Trade, error)
	CancellTrade(*string,*string) error
}