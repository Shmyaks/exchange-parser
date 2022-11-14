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
	p2pURL   string = "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"
	aliasURL string = "https://p2p.binance.com/ru/trade/all-payments/USDT?fiat=RUB"
)

var (
	headers map[string]string = map[string]string{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"origin":     "https://p2p.binance.com",
	}
)

// P2PData struct for binance
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
		"payTypes":       getAliasPayMethod(filter.PayType),
		"countries":      []string{},
		"publisherType":  nil,
		"asset":          string(filter.CryptoCurrency),
		"fiat":           filter.Fiat,
		"tradeType":      filter.TradeType,
		"transAmount":    strconv.Itoa(filter.MinAmount),
	}
	resp, err := d.client.NewRequest().SetBody(body).SetHeaders(headers).Post(p2pURL)
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

		orders = append(
			orders, *models.NewP2POrder(
				info.Adv.Asset,
				info.Adv.FiatUnit,
				info.Advertiser.NickName,
				info.Adv.Price,
				d.marketIDP2P,
				filter.PayType,
			),
		)
	}
	return orders, nil
}

// GetPayMethods get P2P PayMethods for Binance
func (d *P2PData) GetPayMethods() map[models.Fiat][]models.PayMethod {
	mp := make(map[models.Fiat][]models.PayMethod)
	mp[models.RUB] = []models.PayMethod{
		models.Tinkoff, models.Raif, models.Rosbank, models.QIWI, models.YOUMoney, models.Payeer, models.AdvCash,
		models.MTSBank,
	}
	mp[models.USD] = []models.PayMethod{models.AdvCash, models.Payeer, models.Revolut, models.Wise, models.Zelle}
	mp[models.TRY] = []models.PayMethod{models.Ziraat, models.Akbank, models.DenizBank, models.VakifBank}
	return mp
}

func (d *P2PData) GetPayMethodAlias() error {
	return nil
}
