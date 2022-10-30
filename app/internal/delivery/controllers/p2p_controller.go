package controllers

import (
	"sync"
	"time"

	"github.com/Shmyaks/exchange-parser-server/app/config"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
	"github.com/Shmyaks/exchange-parser-server/app/internal/service"
)

// P2PController ... controller of P2P Markets
type P2PController struct {
	service service.P2PService
}

// NewP2PController fabric P2PController
func NewP2PController(service service.P2PService) *P2PController {
	return &P2PController{service: service}
}

// Parse all P2P markets
func (c *P2PController) Parse() {
	var wg sync.WaitGroup
	for {
		wg.Add(len(markets.AllP2PMarket))
		for _, P2Pmarket := range markets.AllP2PMarket {
			go func(market markets.P2PMarket) {
				defer wg.Done()
				c.service.ParseAll(market)
			}(P2Pmarket)
		}

		wg.Wait()
		time.Sleep(config.Env.ParseP2PDelay)
	}
}
