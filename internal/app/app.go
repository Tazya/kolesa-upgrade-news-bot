package app

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"log"
)

func Run(config *config.Config) error {

	s := server.NewServer(config)
	log.Printf("http://localhost:" + config.Port + "/")
	return s.ListenAndServe()

}
