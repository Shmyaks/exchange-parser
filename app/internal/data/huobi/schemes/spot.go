// Package schemes is json schemes of Huobi
package schemes

import "github.com/Shmyaks/exchange-parser-server/app/internal/models"

type allCurencyJSONScheme struct {
	Sc   models.CurencyPairName `json:"sc"`
	Bcdn models.CryptoCurrency  `json:"bcdn"`
	Qcdn models.CryptoCurrency  `json:"qcdn"`
}

// AllCurrencyListJSONScheme json scheme
type AllCurrencyListJSONScheme struct {
	Data []allCurencyJSONScheme `json:"data"`
}

/////////////////////////////////////////

type currencyPriceJSONScheme struct {
	Symbol string  `json:"symbol"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
	Vol    float64 `json:"vol"`
}

// CurrencyPriceListJSONScheme json scheme
type CurrencyPriceListJSONScheme struct {
	Data []currencyPriceJSONScheme `json:"data"`
}
