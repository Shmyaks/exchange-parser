// Package huobi data methods file
package huobi

import (
	"fmt"
	"strings"

	"github.com/Shmyaks/exchange-parser-server/app/internal/data/huobi/schemes"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"

	"github.com/imroc/req/v3"
)

const (
	allURL   string = "https://api.huobi.pro/v2/settings/common/symbols"
	priceURL string = "https://api.huobi.pro/market/tickers"
)

// SPOTData struct
type SPOTData struct {
	client   req.Client
	marketID markets.SPOTMarket
}

// NewSPOTData fabric for Data bybit
func NewSPOTData(client req.Client) *SPOTData {
	return &SPOTData{client: client, marketID: markets.Huobi}
}

// GetMarketID get market id
func (d *SPOTData) GetMarketID() *markets.SPOTMarket {
	return &d.marketID
}

// GetAllPairsAPI get p2p orders from huobi
func (d *SPOTData) GetAllPairsAPI() ([]models.BaseCurrencyPair, error) {
	var pairs []models.BaseCurrencyPair
	var scheme schemes.AllCurrencyListJSONScheme
	resp, err := d.client.NewRequest().Get(allURL)
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
		pairs = append(pairs, *models.NewBaseCurrencyPair(info.Bcdn, info.Bcdn, d.marketID))
	}
	return pairs, nil
}

// GetDetailPairsAPI get currencies price
func (d *SPOTData) GetDetailPairsAPI() ([]models.CurrencyPair, error) {
	currenices := make([]models.CurrencyPair, 0, 10)
	var scheme schemes.CurrencyPriceListJSONScheme
	resp, err := d.client.NewRequest().Get(priceURL)
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
		currenices = append(currenices, *models.NewCurencyPair(models.CurencyPairName(
			strings.ToUpper(info.Symbol)),
			fmt.Sprintf("%v", info.Vol),
			fmt.Sprintf("%v", info.Bid),
			fmt.Sprintf("%v", info.Ask),
			d.marketID))
	}
	return currenices, nil
}
