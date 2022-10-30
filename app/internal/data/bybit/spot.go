// Package bybit data methods file
package bybit

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/bybit/schemes"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"

	"github.com/imroc/req/v3"
)

// url of parsing
const (
	url      string = "https://api.bybit.com/spot/v3/public/symbols"
	priceURL string = "https://api.bybit.com/spot/v3/public/quote/ticker/24hr"
)

// SPOTData sturct for binance
type SPOTData struct {
	client   req.Client
	marketID markets.SPOTMarket
}

// NewSPOTData fabric for SPOTData
func NewSPOTData(client req.Client) *SPOTData {
	return &SPOTData{client: client, marketID: markets.Bybit}
}

// GetMarketID get market id
func (d *SPOTData) GetMarketID() *markets.SPOTMarket {
	return &d.marketID
}

// GetAllPairsAPI get all curencies  from binance
func (d *SPOTData) GetAllPairsAPI() ([]models.BaseCurrencyPair, error) {
	var pairs []models.BaseCurrencyPair
	var scheme schemes.SPOTAllJSONScheme
	resp, err := d.client.NewRequest().Get(url)
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
	for _, info := range scheme.Result.List {
		pairs = append(pairs, *models.NewBaseCurrencyPair(info.BaseCoin, info.QuoteCoin, d.marketID))
	}
	return pairs, nil
}

// GetDetailPairsAPI get detail pairs info of bybit
func (d *SPOTData) GetDetailPairsAPI() ([]models.CurrencyPair, error) {
	pairs := make([]models.CurrencyPair, 0, 10)
	var scheme schemes.SPOTAllDetailJSONScheme

	request := d.client.NewRequest()
	// if filter.CurencyPairsName != nil {
	// 	stringSymbols, _ := jsoniter.Marshal(filter.CurencyPairsName)
	// 	query := map[string]string{
	// 		"symbols": string(stringSymbols),
	// 	}
	// 	request.SetQueryParams(query)
	// }

	resp, err := request.Get(priceURL)
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
	for _, pair := range scheme.Result.List {
		pairs = append(pairs, *models.NewCurencyPair(pair.Symbol, pair.Volume, pair.AskPrice, pair.BidPrice, d.marketID))
	}
	return pairs, nil
}
