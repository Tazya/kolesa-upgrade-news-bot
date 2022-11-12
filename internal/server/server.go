package server

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/http/handlers"
	"net/http"
)

func NewServer() *http.Server {

	port := config.ConfigPort()

	mux := http.NewServeMux()
	handlers.InitRoutes(mux)
	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
