package models

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
)

// CryptoCurrency type for crypto tokens
type CryptoCurrency string

// CurencyPairName type of name currencyPair
type CurencyPairName string

// Fiat type of fiat
type Fiat string

// BaseCurrencyPair type for crypto tokens
type BaseCurrencyPair struct {
	FirstAsset  CryptoCurrency     `json:"firstAsset"`
	SecondAsset CryptoCurrency     `json:"secondAsset"`
	MarketID    markets.SPOTMarket `json:"marketID"`
}

// NewBaseCurrencyPair fabric
func NewBaseCurrencyPair(firstAsset CryptoCurrency, secondAsset CryptoCurrency, marketID markets.SPOTMarket) *BaseCurrencyPair {
	return &BaseCurrencyPair{FirstAsset: firstAsset, SecondAsset: secondAsset, MarketID: marketID}
}

// CurrencyPair struct
type CurrencyPair struct {
	CurencyPairName CurencyPairName    `json:"pair"`
	BuyPrice        string             `json:"buy"`
	Volume          string             `json:"volume"`
	SellPrice       string             `json:"sell"`
	MarketID        markets.SPOTMarket `json:"marketID"`
}

// NewCurencyPair fabric for Pair
func NewCurencyPair(curencyPairName CurencyPairName, volume string, buyPrice string, sellPrice string, marketID markets.SPOTMarket) *CurrencyPair {
	return &CurrencyPair{CurencyPairName: curencyPairName, Volume: volume, BuyPrice: buyPrice, SellPrice: sellPrice, MarketID: marketID}
}

// RUB, USD, TRY, USZ ENUM for Fiat
const (
	RUB Fiat = "RUB"
	USD Fiat = "USD"
	TRY Fiat = "TRY"
	USZ Fiat = "USZ"
)
