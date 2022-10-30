package bybit

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
)

var mapAliasPayMethod = map[models.PayMethod]string{
	models.Tinkoff: "75",
	models.Raif:    "64",
	models.QIWI:    "62",
	models.Rosbank: "185",
}

var mapAliasFilterType = map[filters.TradeType]string{
	filters.Buy:  "1",
	filters.Sell: "0",
}
