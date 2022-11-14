package main

import (
	"flag"
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/app"
	"log"
)

func main() {

	config := &config.Config{}
	port := flag.String("port", "8888", "HTTP port")
	flag.Parse()
	config.Port = *port

	if err := app.Run(config); err != nil {
		log.Fatalf("%v", err)
	}
}
