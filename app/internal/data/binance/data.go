package binance

import "github.com/Shmyaks/exchange-parser-server/app/internal/models"

var mapAliasPayMethod = map[models.PayMethod]string{
	models.Tinkoff:  "TinkoffNew",
	models.Raif:     "RaiffeisenBank",
	models.QIWI:     "QIWI",
	models.Rosbank:  "RosBankNew",
	models.YOUMoney: "YandexMoneyNew",
	models.Payeer:   "Payeer",
	models.AdvCash:  "Advcash",
	models.MTSBank:  "MTSBank",

}

func getAliasPayMethod(payMethod models.PayMethod) []string {
	if val, ok := mapAliasPayMethod[payMethod]; ok {
		return []string{val}
	}
	return []string{string(payMethod)}
}
