package main

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/app"
	"log"
)

func main() {

	config := &config.Config{}

	if err := app.Run(config.Port); err != nil {
		log.Fatalf("%v", err)
	}
}
