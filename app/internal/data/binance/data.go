package binance

import "github.com/Shmyaks/exchange-parser-server/app/internal/models"

var mapAliasPayMethod = map[models.PayMethod]string{
	models.Tinkoff: "TinkoffNew",
	models.Raif:    "RaiffeisenBank",
	models.QIWI:    "QIWI",
	models.Rosbank: "RosBankNew",
}
