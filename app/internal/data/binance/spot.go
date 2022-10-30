// Package binance data methods file
package binance

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/binance/schemes"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"

	"github.com/imroc/req/v3"
)

// url of parsing
const (
	url      string = "https://api.binance.com/api/v3/github.com/Shmyaks/exchange-parser-serverInfo?permissions=SPOT"
	priceURL string = "https://api.binance.com/api/v3/ticker/24hr"
)

// SPOTData sturct for binance
type SPOTData struct {
	client   req.Client
	marketID markets.SPOTMarket
}

// NewSPOTData fabric for BinanceData
func NewSPOTData(client req.Client) *SPOTData {
	return &SPOTData{client: client, marketID: markets.Binance}
}

// GetMarketID get market id
func (d *SPOTData) GetMarketID() *markets.SPOTMarket {
	return &d.marketID
}

// GetAllPairsAPI get all spot pairs of binance
func (d *SPOTData) GetAllPairsAPI() ([]models.BaseCurrencyPair, error) {
	var pairs []models.BaseCurrencyPair
	var scheme schemes.CurrencyJSONScheme
	resp, err := d.client.NewRequest().Get(url)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, nil
	}
	err = resp.Into(&scheme)
	if err != nil {
		return nil, err
	}
	for _, info := range scheme.Symbols {
		pairs = append(pairs, *models.NewBaseCurrencyPair(info.BaseAsset, info.QuoteAsset, d.marketID))
	}
	return pairs, nil
}

// GetDetailPairsAPI get detail pairs info of binance
func (d *SPOTData) GetDetailPairsAPI() ([]models.CurrencyPair, error) {
	pairs := make([]models.CurrencyPair, 0, 10)
	var scheme []schemes.SPOTDetailJSONScheme

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
		return nil, err
	}
	err = resp.Into(&scheme)
	if err != nil {
		return nil, err
	}
	for _, pair := range scheme {
		pairs = append(pairs, *models.NewCurencyPair(pair.Symbol, pair.Volume, pair.AskPrice, pair.BidPrice, d.marketID))
	}
	return pairs, nil
}
