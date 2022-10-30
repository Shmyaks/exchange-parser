package huobi

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
)

var mapAliasFromCurrency = map[models.CryptoCurrency]string{
	"USDT": "2",
	"BTC":  "1",
	"ETH":  "3",
}

var mapAliasToCurrency = map[string]models.CryptoCurrency{
	"2": "USDT",
	"1": "BTC",
	"3": "ETH",
}

var mapAliasFromFiat = map[models.Fiat]string{
	models.RUB: "11",
}

var mapAliasToFiat = map[string]models.Fiat{
	"11": models.RUB,
}

var mapAliasPayMethod = map[models.PayMethod]string{
	models.Tinkoff: "28",
	models.Raif:    "36",
	models.QIWI:    "9",
	models.Rosbank: "358",
}

var mapAliasTradeType = map[filters.TradeType]filters.TradeType{
	filters.Buy:  filters.Sell,
	filters.Sell: filters.Buy,
}
