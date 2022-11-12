package app

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"log"
)

func Run() error {
	port := config.ConfigPort()
	s := server.NewServer()
	log.Printf("http://localhost:" + port + "/")
	return s.ListenAndServe()
}
