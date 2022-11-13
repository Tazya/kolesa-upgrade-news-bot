package main

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/app"
	"log"
	"os"
)

func main() {

	config := &config.Config{}

	for _, arg := range os.Args {
		if arg[:5] == "port=" {
			config.Port = arg[5:]
		}
	}

	if err := app.Run(config); err != nil {
		log.Fatalf("%v", err)
	}
}
