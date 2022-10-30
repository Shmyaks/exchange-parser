// Package schemes is json schemes of Binance
package schemes

import "github.com/Shmyaks/exchange-parser-server/app/internal/models"

type currencyPairJSONScheme struct {
	Symbol     models.CurencyPairName `json:"symbol"`
	BaseAsset  models.CryptoCurrency  `json:"baseAsset"`
	QuoteAsset models.CryptoCurrency  `json:"quoteAsset"`
	Status     string                 `json:"status"`
}

// CurrencyJSONScheme SPot json scheme of Binance: ALL SPOT PAIRS
type CurrencyJSONScheme struct {
	Symbols []currencyPairJSONScheme `json:"symbols"`
}

// SPOTDetailJSONScheme Spot json scheme of Binance: Detail SPOT
type SPOTDetailJSONScheme struct {
	Symbol   models.CurencyPairName `json:"symbol"`
	BidPrice string                 `json:"bidPrice"`
	AskPrice string                 `json:"askPrice"`
	Volume   string                 `json:"volume"`
}
