package main

import (
	"kolesa-upgrade-team/delivery-bot/internal/app"
	"log"
)

func main() {

	if err := app.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
