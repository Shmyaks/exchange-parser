package routes

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal"
	v1 "github.com/Shmyaks/exchange-parser-server/app/internal/delivery/routes/v1"

	"github.com/gofiber/fiber/v2"
)

func InitV1Routes(app *fiber.App, applicationContainer *internal.ApplicationContainer) {

	api := app.Group("/api") // /api
	v1Router := api.Group("/v1")
	v1.SetupP2PRoute(v1Router, applicationContainer.ServiceContainer.P2PService)
	v1.SetupSpotRoute(v1Router, applicationContainer.ServiceContainer.SPOTService)
	// Setup(v1, applicationContainer)
	// SetupCurrencyRoute(v1, *applicationContainer)

}
