package server

import (
	"kolesa-upgrade-team/delivery-bot/http/handlers"
	"net/http"
)

func NewServer(port string) *http.Server {

	mux := http.NewServeMux()
	handlers.InitRoutes(mux)
	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
