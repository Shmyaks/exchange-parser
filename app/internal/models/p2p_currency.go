package models

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
)

// P2POrder struct
type P2POrder struct {
	FirstAsset  CryptoCurrency
	SecondAsset Fiat
	UserName    string
	Price       string
	MarketID    markets.P2PMarket
	PayType     PayMethod
}

// NewP2POrder fabric for P2POrder
func NewP2POrder(firstAsset CryptoCurrency, secondAsset Fiat, userName string, price string, market markets.P2PMarket, payType PayMethod) *P2POrder {
	return &P2POrder{FirstAsset: firstAsset, SecondAsset: secondAsset, UserName: userName, Price: price, MarketID: market, PayType: payType}
}

// P2PPair struct
type P2PPair struct {
	FirstAsset  CryptoCurrency    `json:"firstAsset" dataframe:"FirstAsset,string"`
	SecondAsset Fiat              `json:"secondAsset" dataframe:",string"`
	BuyPrice    string            `json:"buy" dataframe:",string"`
	SellPrice   string            `json:"sell" dataframe:",string"`
	PayType     PayMethod         `json:"payType" dataframe:",string"`
	MarketID    markets.P2PMarket `json:"marketId" dataframe:"marketID,int"`
}

// NewP2PPair fabric for P2PPair
func NewP2PPair(firstAsset CryptoCurrency, secondAsset Fiat, buyPrice string, sellPrice string, market markets.P2PMarket, payType PayMethod) *P2PPair {
	return &P2PPair{FirstAsset: firstAsset, SecondAsset: secondAsset, BuyPrice: buyPrice, SellPrice: sellPrice, MarketID: market, PayType: payType}
}

// GetFullName get full name P2PPair
func (p *P2PPair) GetFullName() string {
	return string(p.FirstAsset) + "/P2P" + string(p.SecondAsset)
}
