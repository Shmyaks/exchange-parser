package schemes

import "github.com/Shmyaks/exchange-parser-server/app/internal/models"

type binanceUserInfoJSONScheme struct {
	NickName string `json:"nickName"`
}

type binanceTradeMethodInfoJSONScheme struct {
	TradeMethodName string `json:"tradeMethodName"`
}

type binanceDetailInfoJSONScheme struct {
	Asset            models.CryptoCurrency              `json:"asset"`
	FiatUnit         models.Fiat                        `json:"fiatUnit"`
	TradableQuantity string                             `json:"tradableQuantity"`
	Price            string                             `json:"price"`
	TradeMethods     []binanceTradeMethodInfoJSONScheme `json:"tradeMethods"`
}

type binanceOrderInfoJSONScheme struct {
	Adv        binanceDetailInfoJSONScheme `json:"adv"`
	Advertiser binanceUserInfoJSONScheme   `json:"advertiser"`
}

// P2PJSONScheme general json scheme for binance
type P2PJSONScheme struct {
	Code string                       `json:"code"`
	Data []binanceOrderInfoJSONScheme `json:"data"`
}
