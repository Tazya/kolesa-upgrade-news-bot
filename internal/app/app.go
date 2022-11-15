package app

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"log"
	"sync"
)

func Run(config *config.Config, wg *sync.WaitGroup) {
	s := server.NewServer(config)
	log.Printf("http://localhost:" + config.Port + "/")

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("%v", err)
	}
}
