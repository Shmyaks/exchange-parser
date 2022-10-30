// Package okx data methods file
package okx

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/okx/schemes"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"

	"github.com/imroc/req/v3"
)

const (
	p2pURL string = "https://www.okx.com/v3/c2c/tradingOrders/books"
)

// P2PData sturct for okx
type P2PData struct {
	client      req.Client
	marketIDP2P markets.P2PMarket
}

// NewP2PData fabric for OkxData
func NewP2PData(client req.Client) *P2PData {
	return &P2PData{client: client, marketIDP2P: markets.OkxP2P}
}

// GetMarketID get p2p market id
func (d *P2PData) GetMarketID() *markets.P2PMarket {
	return &d.marketIDP2P
}

// GetOrdersAPI get P2POrders from okx
func (d *P2PData) GetOrdersAPI(filter filters.P2PFilter) ([]models.P2POrder, error) {
	orders := make([]models.P2POrder, 0, 10)
	var scheme schemes.P2PJSONScheme
	ordersScheme := scheme.Data.Buy

	query := map[string]string{
		"userType":          "all",
		"paymentMethod":     mapAliasPayMethod[filter.PayType],
		"showTrade":         "false",
		"showFollow":        "false",
		"showAlreadyTraded": "false",
		"isAbleFilter":      "false",
		"baseCurrency":      string(filter.CryptoCurrency),
		"quoteCurrency":     string(filter.Fiat),
		"side":              string(mapAliasTradeType[filter.TradeType]),
	}
	resp, err := d.client.NewRequest().SetQueryParams(query).Get(p2pURL)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, nil
	}
	err = resp.Into(&scheme)
	if err != nil {
		return nil, nil
	}
	if mapAliasTradeType[filter.TradeType] == filters.Buy {
		ordersScheme = scheme.Data.Buy
	} else {
		ordersScheme = scheme.Data.Sell
	}
	for _, info := range ordersScheme {
		println(info.Price)
		orders = append(orders, *models.NewP2POrder(
			info.CryptoCurrency,
			info.Fiat,
			info.NickName,
			info.Price,
			d.marketIDP2P,
			filter.PayType))
	}
	return orders, nil
}

// GetPayMethods get P2P PayMethods for Okx
func (d *P2PData) GetPayMethods() map[models.Fiat][]models.PayMethod {
	mp := make(map[models.Fiat][]models.PayMethod)
	mp[models.RUB] = []models.PayMethod{models.Tinkoff, models.Raif, models.Rosbank, models.QIWI}
	return mp
}
