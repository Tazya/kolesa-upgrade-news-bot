package app

import (
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"log"
)

func Run() error {
	s := server.NewServer()
	log.Printf("http://localhost:8888/")
	return s.ListenAndServe()
}
