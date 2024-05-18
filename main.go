package main

import (
	"fmt"
	"log"

	"github.com/lipesalin/go-ecommerce-billing/config"
	"github.com/lipesalin/go-ecommerce-billing/internal/api"
)

func main() {
	// [CONFIG] - Recovery configs
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("[CONFIG] Config file is not loaded properly: %v\n", err)
	}

	fmt.Println("dsadadasdasd")

	// [START SERVER]
	api.StartServer(cfg)
}
