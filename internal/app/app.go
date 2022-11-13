package app

import (
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"log"
)

func Run(port string) error {

	s := server.NewServer(port)
	log.Printf("http://localhost:" + port + "/")
	return s.ListenAndServe()
}
