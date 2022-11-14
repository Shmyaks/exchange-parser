package bybit

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
)

var mapAliasPayMethod = map[models.PayMethod]string{}

var mapAliasFilterType = map[filters.TradeType]string{
	filters.Buy:  "1",
	filters.Sell: "0",
}
