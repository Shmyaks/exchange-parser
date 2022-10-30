package okx // Package bybit data methods file

import (
	"strings"

	"github.com/Shmyaks/exchange-parser-server/app/internal/data/okx/schemes"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"

	"github.com/imroc/req/v3"
)

const (
	priceURL string = "https://www.okx.com/priapi/v5/market/tickers?instType=SPOT"
)

// SPOTData sturct for binance
type SPOTData struct {
	client   req.Client
	marketID markets.SPOTMarket
}

// NewSPOTData fabric for BinanceData
func NewSPOTData(client req.Client) *SPOTData {
	return &SPOTData{client: client, marketID: markets.Okx}
}

// GetMarketID get market id
func (d *SPOTData) GetMarketID() *markets.SPOTMarket {
	return &d.marketID
}

// GetAllPairsAPI get all curencies  from binance
func (d *SPOTData) GetAllPairsAPI() ([]models.BaseCurrencyPair, error) {
	var pairs []models.BaseCurrencyPair
	var scheme schemes.SpotJSONScheme
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
		coins := strings.Split(info.InstIO, "_")
		pairs = append(pairs, *models.NewBaseCurrencyPair(models.CryptoCurrency(coins[0]), models.CryptoCurrency(coins[1]), d.marketID))
	}
	return pairs, nil
}

// GetDetailPairsAPI get currencies price
func (d *SPOTData) GetDetailPairsAPI() ([]models.CurrencyPair, error) {
	pairs := make([]models.CurrencyPair, 0, 10)
	var scheme schemes.SpotJSONScheme
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
	for _, pair := range scheme.Data {
		pairs = append(pairs, *models.NewCurencyPair(models.CurencyPairName(strings.Join(strings.Split(pair.InstIO, "-"), "")), pair.Vol24h, pair.AskPx, pair.BidPx, d.marketID))
	}
	return pairs, nil
}
