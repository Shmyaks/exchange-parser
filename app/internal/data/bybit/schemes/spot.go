package schemes

import "github.com/Shmyaks/exchange-parser-server/app/internal/models"

type curencyJSONScheme struct {
	Name      models.CurencyPairName `json:"name"`
	BaseCoin  models.CryptoCurrency  `json:"baseCoin"`
	QuoteCoin models.CryptoCurrency  `json:"quoteCoin"`
}

type curencyListJSONScheme struct {
	List []curencyJSONScheme `json:"list"`
}

// SPOTAllJSONScheme json scheme
type SPOTAllJSONScheme struct {
	Result curencyListJSONScheme `json:"result"`
}

type currencyDetailInfoJSONSceme struct {
	Symbol   models.CurencyPairName `json:"s"`
	BidPrice string                 `json:"bp"`
	AskPrice string                 `json:"ap"`
	Volume   string                 `json:"v"`
}

type curencyAllDetailListJSONScheme struct {
	List []currencyDetailInfoJSONSceme `json:"list"`
}

// SPOTAllDetailJSONScheme json scheme
type SPOTAllDetailJSONScheme struct {
	Result curencyAllDetailListJSONScheme `json:"result"`
}
