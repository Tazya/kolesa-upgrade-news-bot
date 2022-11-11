package app

import (
	"kolesa-upgrade-team/delivery-bot/internal/server"
	"kolesa-upgrade-team/delivery-bot/pkg/delievery"
	"log"
)

func Run() error {
	handler := delievery.NewHandler()
	s := server.NewServer(*handler)
	log.Printf("http://localhost:8888/")
	return s.ListenAndServe()
}
