package server

import (
	"kolesa-upgrade-team/delivery-bot/http/handlers"
	"net/http"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()
	handlers.InitRoutes(mux)
	return &http.Server{
		Addr:    ":8888",
		Handler: mux,
	}
}
