// Package markets include other market (ENUM)
package markets

// SPOTMarket type of Exchange ENUM
type SPOTMarket uint16

// Binance ... -> ENUM of ExchangeMarket
const (
	Binance SPOTMarket = iota + 1
	Bybit
	Huobi
	Okx
)

// AllSPOTMarkets array of ExchangeMarket ENUM
var AllSPOTMarkets = [4]SPOTMarket{Binance, Bybit, Huobi, Okx}

// GetName get name of ExchangeMarket ENUM
func (m *SPOTMarket) GetName() string {
	var name string
	switch *m {
	case Binance:
		name = "Binance"
	case Bybit:
		name = "Bybit"
	case Huobi:
		name = "Huobi"
	case Okx:
		name = "Okx"
	}
	return name
}
