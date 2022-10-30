// Package filters ...
package filters

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
)

type CurrencyFilter struct {
	CurencyPairsName []models.CurencyPairName
}

func NewCurrencyFilter(curencyPairsName []models.CurencyPairName) *CurrencyFilter {
	return &CurrencyFilter{CurencyPairsName: curencyPairsName}
}

type GetCurrencyByMarket struct {
	MarketID markets.SPOTMarket `validate:"required" json:"marketID"`
}

type GetP2PPairsByMarket struct {
	MarketP2PID markets.P2PMarket `validate:"required" json:"marketP2PID"`
}

type GetCurrencyByPairName struct {
	Pair models.CurencyPairName `validate:"required" json:"pair"`
}
