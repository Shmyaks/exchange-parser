// Package binance data methods file
package binance

import (
	"strconv"

	"github.com/Shmyaks/exchange-parser-server/app/internal/data/binance/schemes"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"

	"github.com/imroc/req/v3"
)

const (
	p2pURL string = "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"
)

// P2PData sturcts for binance
type P2PData struct {
	client      req.Client
	marketIDP2P markets.P2PMarket
}

// NewP2PData fabric for BinanceData
func NewP2PData(client req.Client) *P2PData {
	return &P2PData{client: client, marketIDP2P: markets.BinanceP2P}
}

// GetMarketID get p2p market id
func (d *P2PData) GetMarketID() *markets.P2PMarket {
	return &d.marketIDP2P
}

// GetOrdersAPI get P2POrders from binance
func (d *P2PData) GetOrdersAPI(filter filters.P2PFilter) ([]models.P2POrder, error) {
	orders := make([]models.P2POrder, 0, 10)
	var scheme schemes.P2PJSONScheme
	body := map[string]interface{}{
		"proMerchantAds": false,
		"page":           1,
		"rows":           10,
		"payTypes":       []string{mapAliasPayMethod[filter.PayType]},
		"countries":      []string{},
		"publisherType":  nil,
		"asset":          string(filter.CryptoCurrency),
		"fiat":           filter.Fiat,
		"tradeType":      filter.TradeType,
		"transAmount":    strconv.Itoa(filter.MinAmount),
	}
	resp, err := d.client.NewRequest().SetBody(body).Post(p2pURL)
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
	for _, info := range scheme.Data {

		orders = append(orders, *models.NewP2POrder(
			info.Adv.Asset,
			info.Adv.FiatUnit,
			info.Advertiser.NickName,
			info.Adv.Price,
			d.marketIDP2P,
			filter.PayType))
	}
	return orders, nil
}

// GetPayMethods get P2P PayMethods for Binance
func (d *P2PData) GetPayMethods() map[models.Fiat][]models.PayMethod {
	mp := make(map[models.Fiat][]models.PayMethod)
	mp[models.RUB] = []models.PayMethod{models.Tinkoff, models.Raif, models.Rosbank, models.QIWI}
	return mp
}
package binance

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/binance/schemes"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"

	"github.com/imroc/req/v3"
)

const (
	p2pURL string = "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"
)

// P2PData sturcts for binance
type P2PData struct {
	client      req.Client
	marketIDP2P markets.P2PMarket
}

// NewP2PData fabric for BinanceData
func NewP2PData(client req.Client) *P2PData {
	return &P2PData{client: client, marketIDP2P: markets.BinanceP2P}
}

// GetMarketID get p2p market id
func (d *P2PData) GetMarketID() *markets.P2PMarket {
	return &d.marketIDP2P
}

// GetOrdersAPI get P2POrders from binance
func (d *P2PData) GetOrdersAPI(filter filters.P2PFilter) ([]models.P2POrder, error) {
	orders := make([]models.P2POrder, 0, 10)
	var scheme schemes.P2PJSONScheme
	body := map[string]interface{}{
		"proMerchantAds": false,
		"page":           1,
		"rows":           10,
		"payTypes":       []string{mapAliasPayMethod[filter.PayType]},
		"countries":      []string{},
		"publisherType":  nil,
		"asset":          string(filter.CryptoCurrency),
		"fiat":           filter.Fiat,
		"tradeType":      filter.TradeType,
	}
	resp, err := d.client.NewRequest().SetBody(body).Post(p2pURL)
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
	for _, info := range scheme.Data {
		orders = append(orders, *models.NewP2POrder(
			info.Adv.Asset,
			info.Adv.FiatUnit,
			info.Advertiser.NickName,
			info.Adv.Price,
			d.marketIDP2P,
			filter.PayType))
	}
	return orders, nil
}

// GetPayMethods get P2P PayMethods for Binance
func (d *P2PData) GetPayMethods() map[models.Fiat][]models.PayMethod {
	mp := make(map[models.Fiat][]models.PayMethod)
	mp[models.RUB] = []models.PayMethod{models.Tinkoff, models.Raif, models.Rosbank, models.QIWI}
	return mp
}
