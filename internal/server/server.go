package server

import (
	"kolesa-upgrade-team/delivery-bot/config"
	"kolesa-upgrade-team/delivery-bot/http/handlers"
	"net/http"
)

func NewServer(config *config.Config) *http.Server {
	mux := http.NewServeMux()
	handlers.InitRoutes(mux)
	return &http.Server{
		Addr:    ":" + config.Http.Port,
		Handler: mux,
	}
}
