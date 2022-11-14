// Package data interface file
package data

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
)

// P2P interface for data packages
type P2P interface {
	GetOrdersAPI(filters.P2PFilter) ([]models.P2POrder, error)
	GetPayMethods() map[models.Fiat][]models.PayMethod
	GetMarketID() *markets.P2PMarket
	GetPayMethodAlias() error
}

// SPOT interface for data packages
type SPOT interface {
	GetDetailPairsAPI() ([]models.CurrencyPair, error)
	GetAllPairsAPI() ([]models.BaseCurrencyPair, error)
	GetMarketID() *markets.SPOTMarket
}
