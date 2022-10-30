package controllers

import (
	"sync"
	"time"

	"github.com/Shmyaks/exchange-parser-server/app/config"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
	"github.com/Shmyaks/exchange-parser-server/app/internal/service"
)

// SPOTController controller spot markets
type SPOTController struct {
	currencyService service.SPOTService
}

// NewSPOTController fabric SPOTController
func NewSPOTController(currencyService service.SPOTService) *SPOTController {
	return &SPOTController{currencyService: currencyService}
}

// Parse parse all markets
func (c *SPOTController) Parse() {
	var wg sync.WaitGroup
	for {
		wg.Add(len(markets.AllSPOTMarkets))
		for _, market := range markets.AllSPOTMarkets {
			go func(market markets.SPOTMarket) {
				defer wg.Done()
				c.currencyService.Parse(market)
			}(market)
		}

		wg.Wait()
		time.Sleep(config.Env.ParseSPOTDelay)
	}
}
