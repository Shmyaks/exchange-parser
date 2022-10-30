package okx

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
)

var mapAliasPayMethod = map[models.PayMethod]string{
	models.Tinkoff: "Tinkoff",
	models.Raif:    "Raiffaizen",
	models.QIWI:    "QiWi",
	models.Rosbank: "Rosbank",
}

var mapAliasTradeType = map[filters.TradeType]filters.TradeType{
	filters.Buy:  filters.Sell,
	filters.Sell: filters.Buy,
}
