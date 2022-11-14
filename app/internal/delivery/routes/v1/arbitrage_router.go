package v1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Shmyaks/exchange-parser-server/app/internal/handlers"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/service"
)

// ArbitrageRouter struct
type ArbitrageRouter struct {
	service service.ArbitrageService
}

func SetupArbitrageRoute(api fiber.Router, service service.ArbitrageService) {
	api = api.Group("arbitrage")
	router := ArbitrageRouter{service: service}
	api.Get("/rows", router.getArbitrageRows)
}

// getArbitrageRows godoc
// @Summary     Get currency pair
// @Description Get currency pair from all SPOT markets
// @Tags        Arbitrage
// @Accept      json
// @Produce     json
// @Param       offset query  int true "offset" example(10)
// @Param       limit query  int true "limit" example(10)
// @Success     200  {object} []models.ArbitrageRow
// @Router       /v1/arbitrage/rows [get]
func (r *ArbitrageRouter) getArbitrageRows(c *fiber.Ctx) error {
	var arbitrageFilter filters.ArbitrageFilter
	err := handlers.ValidateQuery(&arbitrageFilter, c)
	if err != nil {
		return err
	}
	rows, err := r.service.Get(arbitrageFilter)
	if err != nil {
		return err
	}
	return c.JSON(rows)
}
