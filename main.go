package main

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/app"
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"log"
)

func main() {

	config := &config.Config{}

	config.Port = server.SetPort()

	if err := app.Run(config); err != nil {
		log.Fatalf("%v", err)
	}
}
