// Package huobi data methods file
package huobi

import (
	"strconv"
	"strings"

	"github.com/Shmyaks/exchange-parser-server/app/internal/data/huobi/schemes"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"

	"github.com/imroc/req/v3"
)

const (
	p2pURL      string = "https://otc-api.trygofast.com/v1/data/trade-market"
	p2pAliasURL string = "https://otc-api.trygofast.com/v1/data/config-list"
)

// P2PData struct
type P2PData struct {
	client      req.Client
	marketIDP2P markets.P2PMarket
}

// NewP2PData fabric for Data bybit
func NewP2PData(client req.Client) *P2PData {
	return &P2PData{client: client, marketIDP2P: markets.HuobiP2P}
}

// GetMarketID get market id
func (d *P2PData) GetMarketID() *markets.P2PMarket {
	return &d.marketIDP2P
}

// GetOrdersAPI get p2p orders from huobi
func (d *P2PData) GetOrdersAPI(filter filters.P2PFilter) ([]models.P2POrder, error) {
	orders := make([]models.P2POrder, 0, 10)
	var scheme schemes.P2PJSONScheme
	query := map[string]string{
		"coinId":       mapAliasFromCurrency[filter.CryptoCurrency],
		"currency":     mapAliasFromFiat[filter.Fiat],
		"tradeType":    string(mapAliasTradeType[filter.TradeType]),
		"currPage":     "1",
		"payMethod":    mapAliasPayMethod[filter.PayType],
		"acceptOrder":  "0",
		"country":      "",
		"blockType":    "general",
		"online":       "1",
		"range":        "0",
		"amount":       strconv.Itoa(filter.MinAmount),
		"isThumbsUp":   "false",
		"isMerchant":   "false",
		"isTraded":     "false",
		"onlyTradable": "false",
		"isFollowed":   "false",
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
	for _, info := range scheme.Data {
		orders = append(
			orders, *models.NewP2POrder(
				mapAliasToCurrency[strconv.Itoa(int(info.CoinID))], mapAliasToFiat[strconv.Itoa(int(info.Currency))],
				info.UserName, info.Price, d.marketIDP2P, filter.PayType,
			),
		)
	}
	return orders, nil
}

// GetPayMethods get P2P PayMethods for bybit
func (d *P2PData) GetPayMethods() map[models.Fiat][]models.PayMethod {
	mp := make(map[models.Fiat][]models.PayMethod)
	mp[models.RUB] = []models.PayMethod{
		models.Tinkoff, models.Raif, models.Rosbank, models.QIWI, models.Payeer, models.AdvCash, models.YOUMoney,
	}
	mp[models.USD] = []models.PayMethod{models.AdvCash, models.Payeer, models.Zelle, models.Wise}
	mp[models.TRY] = []models.PayMethod{models.DenizBank, models.VakifBank, models.Akbank}
	return mp
}

func (d *P2PData) GetPayMethodAlias() error {
	var scheme schemes.AliasJSONScheme
	mp := make(map[models.PayMethod]string)
	query := map[string]string{
		"type": "currency,marketQuery,pay,allCountry,coin",
	}
	resp, err := d.client.NewRequest().SetQueryParams(query).Get(p2pAliasURL)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return nil
	}
	err = resp.Into(&scheme)
	for _, info := range scheme.Data.PayMethod {
		mp[models.PayMethod(
			strings.ToUpper(
				strings.ReplaceAll(
					info.Name, " ", "",
				),
			),
		)] = strconv.Itoa(info.PayMethodID)
	}
	mapAliasPayMethod = mp
	return nil
}
