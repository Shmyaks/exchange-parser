package v1

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/handlers"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
	"github.com/Shmyaks/exchange-parser-server/app/internal/service"

	"github.com/gofiber/fiber/v2"
)

// SPOTRouter ...
type SPOTRouter struct {
	SPOTservice service.SPOTService
}

// SetupSpotRoute init SPOTRouters api
func SetupSpotRoute(api fiber.Router, SPOTService service.SPOTService) {
	api = api.Group("spot")
	router := SPOTRouter{SPOTservice: SPOTService}
	api.Get("/pair", router.getCurrencyAPI)
	api.Get("/pairs/:marketID", router.getAllPairs)
	api.Get("/markets/all", router.getAllMarkets)
}

// getAllMarkets godoc
// @Summary     Get all
// @Description Get all SPOT markets
// @Tags        SPOT
// @Accept      json
// @Produce     json
// @Success     200 {object} []markets.Market
// @Router      /v1/spot/markets/all [get]
func (r *SPOTRouter) getAllMarkets(c *fiber.Ctx) error {
	spotMarkets := make([]markets.Market, 0, len(markets.AllSPOTMarkets))
	for _, val := range markets.AllSPOTMarkets {
		spotMarkets = append(spotMarkets, *markets.NewMarket(uint16(val), val.GetName()))
	}
	return c.JSON(spotMarkets)
}

// getCurrencyAPI godoc
// @Summary     Get all
// @Description Get all SPOT pairs by market id
// @Tags        SPOT
// @Accept      json
// @Produce     json
// @Param       marketID path     string true "SPOT market id"
// @Success     200      {object} []models.CurrencyPair
// @Router      /v1/spot/pairs/{marketID} [get]
func (r *SPOTRouter) getAllPairs(c *fiber.Ctx) error {
	param := new(filters.GetCurrencyByMarket)
	err := handlers.ValidateParam(param, c)
	if err != nil {
		return err
	}
	pairs := r.SPOTservice.GetAll(param.MarketID)
	return c.JSON(pairs)
}

// getAllPairs godoc
// @Summary     Get currency pair
// @Description Get currency pair from all SPOT markets
// @Tags        SPOT
// @Accept      json
// @Produce     json
// @Param       pair query    string true "Name of SPOT pair" example(BTCUSDT)
// @Success     200  {object} []models.P2PPair
// @Router      /v1/spot/pair [get]
func (r *SPOTRouter) getCurrencyAPI(c *fiber.Ctx) error {
	query := new(filters.GetCurrencyByPairName)
	err := handlers.ValidateQuery(query, c)
	if err != nil {
		return err
	}
	pairs := make([]models.CurrencyPair, 0, len(markets.AllSPOTMarkets))
	for _, market := range markets.AllSPOTMarkets {
		pair := r.SPOTservice.Get(market, models.CurencyPairName(query.Pair))
		pairs = append(pairs, *pair)
	}
	return c.JSON(pairs)
}
