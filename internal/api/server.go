package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/lipesalin/go-ecommerce-billing/config"

)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	// [MOUNT URL] - example: localhost:3030
	serverListen := fmt.Sprintf("%s:%s", config.ServerUrl, config.ServerPort)

	// [LISTEN] - Listen server
	app.Listen(serverListen)
}
