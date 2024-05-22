package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lipesalin/go-ecommerce-billing/config"
	"github.com/lipesalin/go-ecommerce-billing/internal/api/rest"
	"github.com/lipesalin/go-ecommerce-billing/internal/api/rest/handlers"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	restHandler := &rest.RestHandler{
		App: app,
	}

	// [ROUTES]
	setupRoutes(restHandler)

	// [MOUNT URL] - example: localhost:3030
	serverListen := fmt.Sprintf("%s:%s", config.ServerUrl, config.ServerPort)

	// [LISTEN] - Listen server
	app.Listen(serverListen)
}

func setupRoutes(rh *rest.RestHandler) {
	// user
	handlers.SetupUserRoutes(rh)
	// transactions
	// catalog
}
