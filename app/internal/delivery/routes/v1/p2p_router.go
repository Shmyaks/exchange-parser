// Package v1 ...
package v1

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/handlers"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
	"github.com/Shmyaks/exchange-parser-server/app/internal/service"

	"github.com/gofiber/fiber/v2"
)

// P2PRouter struct
type P2PRouter struct {
	service service.P2PService
}

// SetupP2PRoute P2P routes
func SetupP2PRoute(api fiber.Router, service service.P2PService) {
	api = api.Group("p2p")
	router := P2PRouter{service: service}
	api.Get("/pair", router.getCurrencyAPI)
	api.Get("/pairs/:marketP2PID", router.getAllPairs)
	api.Get("/markets/all", router.getAllMarkets)
}

// getAllMarkets godoc
// @Summary     Get all P2P markets
// @Description Get all P2P markets
// @Tags        P2P
// @Accept      json
// @Produce     json
// @Success     200 {object} []markets.Market
// @Router      /v1/p2p/markets/all [get]
func (r *P2PRouter) getAllMarkets(c *fiber.Ctx) error {
	p2pMarkets := make([]markets.Market, 0, len(markets.AllP2PMarket))
	for _, val := range markets.AllP2PMarket {
		p2pMarkets = append(p2pMarkets, *markets.NewMarket(uint16(val), val.GetName()))
	}
	return c.JSON(p2pMarkets)
}

// getAllPairs godoc
// @Summary     Get all
// @Description Get all orders by market id
// @Tags        P2P
// @Accept      json
// @Produce     json
// @Param       marketP2PID path     string true "Name of SPOT pair"
// @Success     200         {object} []models.P2PPair
// @Router      /v1/p2p/pairs/{marketP2PID} [get]
func (r *P2PRouter) getAllPairs(c *fiber.Ctx) error {
	param := new(filters.GetP2PPairsByMarket)
	err := handlers.ValidateParam(param, c)
	if err != nil {
		return err
	}
	pairs := r.service.GetAll(param.MarketP2PID)
	return c.JSON(pairs)
}

// getCurrencyAPI godoc
// @Summary     Get
// @Description Get p2p orders from all markets
// @Tags        P2P
// @Accept      json
// @Produce     json
// @Param       pair query    string true "Name of P2P pair" example(USDT/P2PRUB)
// @Success     200  {object} []models.P2PPair
// @Router      /v1/p2p/pair [get]
func (r *P2PRouter) getCurrencyAPI(c *fiber.Ctx) error {
	query := new(filters.GetCurrencyByPairName)
	err := handlers.ValidateQuery(query, c)
	if err != nil {
		return err
	}
	pairs := make([]models.P2PPair, 0, 100)
	for _, market := range markets.AllP2PMarket {
		pair := r.service.Get(market, query.Pair)
		pairs = append(pairs, pair...)
	}
	return c.JSON(pairs)
}
